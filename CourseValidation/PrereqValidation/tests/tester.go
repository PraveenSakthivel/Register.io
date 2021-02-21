package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	prereqInterface "registerio/cv/preqreq/protobuf"

	"google.golang.org/grpc"
)

type Case struct {
	Name      string           `json:"Case"`
	History   map[string]int32 `json:"history"`
	Requested []string         `json:"requested"`
	Results   map[string]bool  `json:"results"`
}

func buildCases() ([]Case, error) {
	jsonFile, err := os.Open("cases.json")
	if err != nil {
		fmt.Println("Error Opening Cases.json: ", err)
		return nil, err
	}
	defer jsonFile.Close()

	var cases []Case
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error Reading Test Cases: ", err)
		return nil, err
	}

	err = json.Unmarshal(byteValue, &cases)
	if err != nil {
		fmt.Println("Error Reading Parsing Json: ", err)
		return nil, err
	}

	return cases, nil
}

func main() {
	cases, err := buildCases()
	if err != nil {
		return
	}
	casesPassed := 0
	fmt.Printf("Running %d Cases\n-------------------\n", len(cases))
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("ERROR: Could not connect to server: ", err)
		return
	}

	defer conn.Close()
	server := prereqInterface.NewPrereqValidationClient(conn)
	for _, test := range cases {
		fmt.Println("Trying Case: ", test.Name)
		request := prereqInterface.PrereqRequest{ClassHistory: test.History, Indices: test.Requested}

		response, err := server.CheckPrereqs(context.Background(), &request)

		if err != nil {
			fmt.Printf("Error Making Request: %s\n\n", err.Error())
			continue
		}

		fmt.Println("Results:\nClass\tExpec\tResult")
		results := response.Results
		expected := test.Results
		for class, out := range results {
			fmt.Println(class, "\t", expected[class], "\t", out)
		}

		pass := reflect.DeepEqual(results, expected)
		if pass {
			fmt.Println("Test Passed")
			casesPassed++
		} else {
			fmt.Println("Test Failed")
		}
		fmt.Print("\n\n")
	}

	fmt.Printf("%d/%d cases passed. %f accuracy", casesPassed, len(cases), float64(casesPassed)/float64(len(cases)))
}
