package rvInterface

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"context"
)

//RV Server struct, contains dataset and debug field
type Server struct {
	UnimplementedRegistrationValidationServer
	students map[string]int
	debug    bool
}

//DB Info
const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

func (s *Server) dprint(msg ...interface{}) {
	if s.debug {
		log.Println(msg...)
	}
}

// Retrieve list of all students from Database
// TODO: Retrieve endpoint securely
func (s *Server) retrieveData() map[string]int {
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

	s.dprint("OK: Successfully Pulled Data from DB")
	return students
}

//Downloads information and creates new server instance
func NewServer() *Server {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	s := &Server{students: make(map[string]int), debug: *debugPrnt}
	s.students = s.retrieveData()
	return s
}

//Checks whether user is elgibile to register
func (s *Server) CheckRegVal(ctx context.Context, student *Student) (*Response, error) {
	resp := Response{
		Eligible: false,
		Error:    "",
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

	s.dprint("OK: Request from with NetID: ", student.NetId)
	return &resp, nil
}
