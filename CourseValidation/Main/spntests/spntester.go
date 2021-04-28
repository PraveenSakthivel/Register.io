package main

import (
	"context"
	"fmt"
	cvInterface "registerio/cv/main/protobuf"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("ERROR: Could not connect to server: ", err)
		return
	}

	defer conn.Close()
	server := cvInterface.NewCourseValidationClient(conn)
	fmt.Println("Trying Case: Default")
	test := cvInterface.SPNRequest{Spn: "11704", Index: "02345", Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2MTQ1NjA1MDUsImV4cCI6MTY0NjA5NjUwNSwibmFtZSI6InBzOTMxIiwiY2xhc3NIaXN0b3J5Ijp7IjE5MjoxMDEiOjQwMCwiMTkyOjEwMiI6NDAwfX0.OiyDhMIu0xlgj2036vGH4JCDkELRWktvian3pdeOFSU"}
	response, err := server.AddSPN(context.Background(), &test)

	if err != nil {
		fmt.Printf("Error Making Request: %s\n\n", err.Error())
		return
	}
	if response.Valid {
		fmt.Println("Test Passed")
	} else {
		fmt.Println("Test Failed")
	}
	fmt.Print("\n\n")
}
