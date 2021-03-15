package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	rvInterface "registerio/rv/protobuf"
	secret "registerio/rv/secrets"
	"time"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

var debug = false

//DB Info
type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

type Server struct {
	rvInterface.UnimplementedRegistrationValidationServer
	students map[string]int
	debug    bool
	db       *DB
}

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func buildDB() (*DB, error) {
	dbstring, err := secret.GetTokenSecret("prod/DB")
	if err != nil {
		return nil, err
	}
	retval := DB{}
	err = json.Unmarshal([]byte(dbstring), &retval)
	if err != nil {
		return nil, err
	}
	return &retval, nil
}

// Retrieve list of all students from Database
// TODO: Retrieve endpoint securely
func (s *DB) retrieveData() map[string]int {
	students := make(map[string]int)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
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

//Pulls Info from DB and creates new Server Struct
func NewServer() *Server {
	db, err := buildDB()
	if err != nil {
		log.Fatal("Error building database: ", err)
	}
	students := db.retrieveData()
	s := &Server{students: students, debug: debug, db: db}
	return s
}

//Checks whether user is elgibile to register
func (s *Server) CheckRegVal(ctx context.Context, student *rvInterface.Student) (*rvInterface.Response, error) {
	resp := rvInterface.Response{
		Eligible: false,
	}
	// Check to see if student is eligible
	if dateInt, ok := s.students[student.NetId]; ok {
		date := time.Unix(int64(dateInt), 0)
		if time.Now().After(date) {
			resp.Eligible = true
		}
	} else {
		log.Println("WARNING: Unidentifiable NetID ", student.NetId)
	}

	dprint("OK: Request with NetID: ", student.NetId)
	return &resp, nil
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	debug = *debugPrnt

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}

	s := NewServer()
	grpcServer := grpc.NewServer()
	rvInterface.RegisterRegistrationValidationServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
