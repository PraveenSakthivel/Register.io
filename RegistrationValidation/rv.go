package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	rvInterface "registerio/rv/protobuf"
	"time"

	_ "github.com/lib/pq"

	proto "github.com/golang/protobuf/proto"
)

//Global list of all students with format netId:EligbilityTimestamp
var students map[string]int

//DB Info
const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

// Retrieve list of all students from Database
// TODO: Retrieve endpoint securely
func retrieveData() {
	students = make(map[string]int)

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

}

//Send response back to client
func sendResp(resp *rvInterface.Response, w http.ResponseWriter) {
	data, err := proto.Marshal(resp)
	if err != nil {
		log.Println("Encoding error: ", err)
		resp.Error += "\n" + err.Error()
	}
	w.Write(data)
}

// Check whether student is eligbile to register
func checkRegVal(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	resp := &rvInterface.Response{
		Eligible: false,
		Error:    "",
	}

	// Read request body
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Reading Body: ", err)
		resp.Error = err.Error()
		sendResp(resp, w)

	}

	// Decode body into Student struct
	student := &rvInterface.Student{}
	if err := proto.Unmarshal(input, student); err != nil {
		log.Println("Decoding Error :", err)
		resp.Error = err.Error()
		sendResp(resp, w)
	}

	// Check to see if student is eligible
	if dateInt, ok := students[student.NetId]; ok {
		date := time.Unix(int64(dateInt), 0)
		if time.Now().After(date) {
			resp.Eligible = true
		}
	}

	sendResp(resp, w)
}

func main() {
	retrieveData()
	http.HandleFunc("/", checkRegVal)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
