package main

import (
	"database/sql"
	"fmt"
	"flag"
	"log"
	"os"
	"net"
	dbRequests "registerio/db/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/lib/pq"
)


const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

type Env struct {
	db *sql.DB
}



func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	var debug = *debugPrnt
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Unable to open connection to Postgres DB: ", err)
		os.Exit(3)
	}

	env := &Env{db: db}

	log.Println("Successfully connected to Postgres Database")
	log.Printf("Connection Details: host=%s port=%d user=%s "+
	"dbname=%s sslmode=disable\n", host, port, user, dbname)

	var port string
	var ok bool
	port, ok = os.LookupEnv("PORT")
	if ok {
		log.Printf("PORT: %s\n", port)
	} else {
		port = "9000"
		log.Println("PORT not defined, going with default (9000)")
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error: Failed to listen")
	}

	s := dbRequests.Server{Db:env.db, Debug:debug}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	dbRequests.RegisterDatabseWrapperServer(grpcServer, &s)

	log.Println("gRPC server started at ", port)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve")
	}

}