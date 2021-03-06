package prereqInterface

import (
	context "context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"time"

	cvInterface "registerio/cv/main/protobuf"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const prereqEndpoint = "prereq.registerio.co:8080"

func CheckPrereqs(classHistory map[string]int32, indices []string, cases map[int32]bool) (*PrereqResponse, error) {
	prereq := PrereqRequest{ClassHistory: classHistory, Indices: indices, Cases: cases}
	var preResult *PrereqResponse
	var conn *grpc.ClientConn
	var err = errors.New("TMP")
	start := time.Now()
	certPool, _ := x509.SystemCertPool()
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            certPool,
	}
	for err != nil {
		conn, err = grpc.Dial(prereqEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(config)))
		defer conn.Close()
		if err != nil {
			log.Println("ERROR: Unable to connect to Prereq Endpoint: ", err)
			if time.Now().Sub(start).Milliseconds() > 1000 {
				log.Println("ERROR: Prereq Timeout")
				return nil, errors.New("Timeout")
			}
			continue
		}
		server := NewPrereqValidationClient(conn)
		preResult, err = server.CheckPrereqs(context.Background(), &prereq)
		if err != nil {
			log.Println("ERROR: Error from Prereq Endpoint: ", err)
			return nil, err
		}
	}
	return preResult, nil
}

func EvalPrereqResults(preResult *PrereqResponse, results *map[string]cvInterface.ResultClass, classes []*cvInterface.ClassOperations, debug bool) []*cvInterface.ClassOperations {
	//debug print
	dprint := func(msg ...interface{}) {
		if debug {
			log.Println(msg...)
		}
	}

	eligibleClasses := make(map[string]bool)
	resultMap := *results
	var retVal []*cvInterface.ClassOperations

	if preResult.InvalidIndices != nil {
		for _, index := range preResult.InvalidIndices {
			dprint("Invalid Index: ", index)
			resultMap[index] = cvInterface.ResultClass_INVALID
		}
	}

	for index, status := range preResult.Results {
		if status == false {
			if resultMap[index] != cvInterface.ResultClass_INVALID {
				dprint("Prereq Failed: ", index)
				resultMap[index] = cvInterface.ResultClass_PREREQ
			}
		} else {
			dprint("Class OK: ", index)
			eligibleClasses[index] = true
			resultMap[index] = cvInterface.ResultClass_OK
		}
	}

	for _, class := range classes {
		if _, ok := eligibleClasses[class.Index]; ok {
			retVal = append(retVal, class)
		}
	}
	return retVal
}
