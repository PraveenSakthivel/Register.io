package main

import (
	"context"
	"flag"
	"log"
	"net"
	cvInterface "registerio/cv/main/protobuf"

	"google.golang.org/grpc"
)

var debug = false

//Server instance
type Server struct {
	cvInterface.UnimplementedCourseValidationServer
}

//Student representation for handling request
type Student struct {
	netID        string
	classHistory map[string]string
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
	return Student{netID: "ps931", classHistory: {"198:101": "A", "198:102", "A"}}, nil
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
		return &response, nil
	}
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
