package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	cvInterface "registerio/cv/main/protobuf"

	"google.golang.org/grpc"
)

type ClassRequest struct {
	Index     string `json:"Index"`
	Operation string `json:"Operation"`
}

type Req struct {
	Name    string            `json:"Case"`
	Token   string            `json:"Token"`
	Classes []ClassRequest    `json:"Classes"`
	Results map[string]string `json:"Results"`
}

type Case struct {
	Name    string
	Classes *cvInterface.RegistrationRequest
	Results map[string]string
}

func buildCases() ([]Case, error) {
	var retVal []Case
	var reqs []Req
	jsonFile, err := os.Open("cases.json")
	if err != nil {
		fmt.Println("Error Opening Cases.json: ", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error Reading Test Cases: ", err)
		return nil, err
	}

	err = json.Unmarshal(byteValue, &reqs)
	if err != nil {
		fmt.Println("Error Reading Parsing Json: ", err)
		return nil, err
	}

	for _, test := range reqs {
		var ops []*cvInterface.ClassOperations
		for _, op := range test.Classes {
			var classOp cvInterface.ClassOperations
			if op.Operation == "add" {
				classOp = cvInterface.ClassOperations{Index: op.Index, Op: cvInterface.ReqOp_ADD}
			} else {
				classOp = cvInterface.ClassOperations{Index: op.Index, Op: cvInterface.ReqOp_DROP}
			}
			ops = append(ops, &classOp)
		}
		req := cvInterface.RegistrationRequest{Token: test.Token, Classes: ops}
		newCase := Case{Name: test.Name, Classes: &req, Results: test.Results}
		retVal = append(retVal, newCase)
	}

	return retVal, nil
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
	server := cvInterface.NewCourseValidationClient(conn)
	for _, test := range cases {
		fmt.Println("Trying Case: ", test.Name)
		response, err := server.ChangeRegistration(context.Background(), test.Classes)

		if err != nil {
			fmt.Printf("Error Making Request: %s\n\n", err.Error())
			continue
		}
		fmt.Println("Results:\nIndex\tExpec\t\tResult")
		var passed = true
		expected := test.Results
		for index, actual := range response.Results {
			actualString := actual.String()
			if actualString != expected[index] {
				passed = false
			}
			fmt.Println(index, "\t", expected[index], "\t\t", actualString)
		}
		if passed {
			casesPassed++
			fmt.Println("Test Passed")
		} else {
			fmt.Println("Test Failed")
		}
		fmt.Print("\n\n")
	}

	fmt.Printf("%d/%d cases passed. %f accuracy", casesPassed, len(cases), float64(casesPassed)/float64(len(cases)))
}
