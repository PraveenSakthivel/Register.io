package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	rvInterface "registerio/rv/validation"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

var debug = false

//DB Info
const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

// Retrieve list of all students from Database
// TODO: Retrieve endpoint securely
func retrieveData() map[string]int {
	students := make(map[string]int)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}

	rows, err := db.Query("SELECT * FROM \"registration dates\"")
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}
	defer rows.Close()
	for rows.Next() {
		var netid string
		var time int
		err = rows.Scan(&netid, &time)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			os.Exit(3)
		}
		students[netid] = time
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		os.Exit(3)
	}

	dprint("OK: Successfully Pulled Data from DB")
	return students
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	debug = *debugPrnt
	//students := retrieveData()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}

	s := rvInterface.NewServer()
	grpcServer := grpc.NewServer()
	rvInterface.RegisterRegistrationValidationServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
