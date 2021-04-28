package main

import (
	"flag"
	"log"
	"os"
	"net"
	dbRequests "registerio/db/pb"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
	data "registerio/db/database"
)

type Env struct {
	db *data.DB
	tokenSecret string
}

type token struct {
	TokenSecret string
}

func setup() (*Env){
	var err error
	db, err := data.BuildDB()
	if err != nil {
		log.Fatal("ERROR Unable to build DB: ", err)
		return nil
	}

	tokenSecret, err := data.GetTokenSecret()
	if err != nil {
		log.Fatal("ERROR Unable to get token secret: ", err)
		return nil
	}

	return &Env{db: db, tokenSecret: tokenSecret.TokenSecret}
	
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	var debug = *debugPrnt
	
	env := setup()

	if env == nil {
		return 
	}

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

	s := dbRequests.Server{Db:env.db, Debug:debug, TokenSecret: env.tokenSecret}

	grpcServer := grpc.NewServer()

	dbRequests.RegisterDatabaseWrapperServer(grpcServer, &s)

	log.Println("gRPC server started at ", port)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve")
	}

}