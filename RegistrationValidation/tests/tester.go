package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	rvInterface "registerio/rv/validation"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
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
	for netID, result := range cases {
		fmt.Println("Trying Case: ", netID)
		student := &rvInterface.Student{NetId: netID}
		data, err := proto.Marshal(student)
		if err != nil {
			fmt.Printf("Error Encoding Test Case: %s\n\n", err.Error())
			continue
		}

		resp, err := http.Post("http://localhost:8080/", "application/Protobuf", bytes.NewBuffer(data))

		if err != nil {
			fmt.Printf("Error Making Request: %s\n\n", err.Error())
			continue
		}

		input, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error Reading Body: %s\n\n", err)
			continue
		}

		response := &rvInterface.Response{}
		if err := proto.Unmarshal(input, response); err != nil {
			fmt.Printf("Error Decoding Response: %s\n\n", err)
			continue
		}

		if response.Error != "" {
			fmt.Printf("Error on Server Side: %s\n\n", response.Error)
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
