package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	data "registerio/cv/preqreq/database"
	prereqInterface "registerio/cv/preqreq/protobuf"

	"google.golang.org/grpc"
)

var debug = false

//Server Instance
type Server struct {
	prereqInterface.UnimplementedPrereqValidationServer
	indexLookup map[string]string //Matches Index to Course Number
	//Prereq Store. Prereqs represented by a 2d array. Key is the class number. Each row is a prereq, each columns is a class that fulfills it (coreqs)
	prereqs map[string][][]data.Prereq
}

//debug print
func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

//Initialize any Server Variables
func NewServer() *Server {
	prereqs, err := data.GetPrereqs()
	if err != nil {
		os.Exit(3)
	}
	dprint("Prereqs: ", prereqs)
	lookup, err := data.GetLookups()
	if err != nil {
		os.Exit(3)
	}
	dprint("Lookup: ", lookup)
	s := &Server{indexLookup: lookup, prereqs: prereqs}
	return s
}

//Check if student is eligible to register for class
func (s *Server) CheckPrereqs(ctx context.Context, req *prereqInterface.PrereqRequest) (*prereqInterface.PrereqResponse, error) {
	response := prereqInterface.PrereqResponse{Results: make(map[string]bool), InvalidIndices: nil}
	history := req.ClassHistory
	dprint("Received Request")
	//Loop through all requested class
	for _, index := range req.Indices {
		dprint("Request Index: ", index)
		if _, ok := s.indexLookup[index]; !ok {
			response.Results[index] = false
			response.InvalidIndices = append(response.InvalidIndices, index)
			continue
		}
		class := s.indexLookup[index]
		if _, ok := s.prereqs[class]; !ok {
			response.Results[index] = true
			continue
		}
		dprint("Request class: ", class)
		reqs := s.prereqs[class]
		eligible := false
		//Loop through all preq "sets"
		for _, req := range reqs {
			eligible = false
			//Loop through all possibilities in preq set and check if at least one matches
			for _, coreq := range req {
				if grade, ok := history[coreq.Class]; ok {
					if int(grade) >= coreq.Grade {
						eligible = true
						break
					}
				}
			}
			//if eligible has not been set to true after looping through all possibiltiies, not eligible to register
			if !eligible {
				response.Results[index] = false
				break
			}
		}
		//If after looping through all requirements, eligible is still true. Eligible to register
		if eligible {
			response.Results[index] = true
		}
	}
	return &response, nil
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	debug = *debugPrnt

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Failed to listen on port 8081: ", err)
	}

	s := NewServer()
	grpcServer := grpc.NewServer()
	prereqInterface.RegisterPrereqValidationServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8081: ", err)
	}
}
