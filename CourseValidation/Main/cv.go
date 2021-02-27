package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	prereqInterface "registerio/cv/main/prereq"
	cvInterface "registerio/cv/main/protobuf"
	"time"

	"google.golang.org/grpc"
)

var debug = false

const PREREQ_ENDPOINT = ":8081"

//Server instance
type Server struct {
	cvInterface.UnimplementedCourseValidationServer
}

//Student representation for handling request
type Student struct {
	netID        string
	classHistory map[string]int32
}

//debug print
func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

//Initialize any Server Variables
func NewServer() *Server {
	s := &Server{}
	return s
}

//Dummy function, replace with imported library later
func parseJWT(token string) (Student, error) {
	return Student{netID: "ps931", classHistory: {"198:101": 40, "198:102", 40}}, nil
}

func checkPrereqs(classHistory map[string]int32, indices []string) (*prereqInterface.PrereqResponse, error) {
	prereq := prereqInterface.PrereqRequest{ClassHistory: classHistory, Indices: indices}
	var preResult *prereqInterface.PrereqResponse
	var conn *grpc.ClientConn
	var err error
	start := time.Now()
	for err != nil {
		conn, err = grpc.Dial(PREREQ_ENDPOINT, grpc.WithInsecure())
		if err != nil {
			log.Panic("ERROR: Unable to connect to Prereq Endpoint: ", err)
			continue
		}
		server := prereqInterface.NewPrereqValidationClient(conn)
		ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
		preResult, err = server.CheckPrereqs(ctx, &prereq)
		if err != nil {
			log.Panic("ERROR: Unable to connect to Prereq Endpoint: ", err)
		} else if time.Now().Sub(start) > 1000 {
			return nil, errors.New("Timeout")
		}
	}
	conn.Close()
	return preResult, nil
}

func evalPrereqResults(preResult *prereqInterface.PrereqResponse, results *map[string]cvInterface.ResultClass) []string {
	var eligibleClasses []string
	resultMap := *results

	for _, index := range preResult.InvalidIndices {
		resultMap[index] = cvInterface.ResultClass_INVALID
	}

	for index, status := range preResult.Results {
		if status == false {
			resultMap[index] = cvInterface.ResultClass_PREREQ
		} else {
			eligibleClasses = append(eligibleClasses, index)
			resultMap[index] = cvInterface.ResultClass_OK
		}
	}
	return eligibleClasses
}

func (s *Server) ChangeRegistration(ctx context.Context, req *cvInterface.RegistrationRequest) (*cvInterface.RegistrationResponse, error) {
	results := make(map[string]cvInterface.ResultClass)
	response := cvInterface.RegistrationResponse{Results: results}

	indices := req.Indices
	for _, index := range indices {
		response.Results[index] = cvInterface.ResultClass_ERROR
	}

	student, err := parseJWT(req.Token)
	if err != nil {
		log.Panic("ERROR Parsing JWT")
		return &response, nil
	}

	preResult, err := checkPrereqs(student.classHistory, indices)
	if err != nil {
		log.Panic("ERROR Timed out trying to connect to prereq endpoint")
		return &response, nil
	}

	eligibleClasses := evalPrereqResults(preResult, &results)

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
	cvInterface.RegisterCourseValidationServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
