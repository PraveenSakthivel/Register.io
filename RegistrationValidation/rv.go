package main

import (
	"io/ioutil"
	"log"
	"net/http"
	rvInterface "registerio/rv/protobuf"
	"time"

	proto "github.com/golang/protobuf/proto"
)

//Global list of all students with format netId:EligbilityTimestamp
var students map[string]int

// Retrieve list of all students from Database
// TODO: Implement API Call to Database and construct Student Objects from returned Data
func retrieveData() {
	students = map[string]int{"ps931": 2309483, "mg123": 1712129765}
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
