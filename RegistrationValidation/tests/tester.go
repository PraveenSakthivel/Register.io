package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	rvInterface "registerio/rv/protobuf"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func buildCases() map[string]bool {
	f, err := os.Open("cases.txt")
	retval := make(map[string]bool)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		slices := strings.Split(scanner.Text(), "|")
		retval[slices[0]], _ = strconv.ParseBool(slices[1])
	}

	return retval
}

func main() {
	cases := buildCases()
	fmt.Printf("Running %d Cases\n-------------------\n", len(cases))
	casesPassed := 0
	var conn *grpc.ClientConn
	certPool, err := x509.SystemCertPool()
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            certPool,
	}
	conn, err = grpc.Dial("rv.registerio.co:8080", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		fmt.Printf("ERROR: Could not connect to server ", err)
	}

	defer conn.Close()
	server := rvInterface.NewRegistrationValidationClient(conn)

	for token, result := range cases {
		fmt.Println("Trying Case: ", token)
		student := rvInterface.Student{Token: token}

		response, err := server.CheckRegVal(context.Background(), &student)

		if err != nil {
			fmt.Printf("Error Making Request: %s\n\n", err.Error())
			continue
		}

		if response.Eligible == result {
			fmt.Printf("Passed\n\n")
			casesPassed++
		} else {
			fmt.Printf("Failed. Got %t, Expected %t\n\n", response.Eligible, result)
		}
	}
	fmt.Printf("%d/%d cases passed. %f accuracy", casesPassed, len(cases), float64(casesPassed)/float64(len(cases)))
}
