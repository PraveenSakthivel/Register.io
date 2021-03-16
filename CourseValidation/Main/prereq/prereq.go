package prereqInterface

import (
	context "context"
	"errors"
	"log"
	"time"

	cvInterface "registerio/cv/main/protobuf"

	grpc "google.golang.org/grpc"
)

const prereqEndpoint = "3.228.3.112:8080"

func CheckPrereqs(classHistory map[string]int32, indices []string, cases map[int32]bool) (*PrereqResponse, error) {
	prereq := PrereqRequest{ClassHistory: classHistory, Indices: indices, Cases: cases}
	var preResult *PrereqResponse
	var conn *grpc.ClientConn
	var err = errors.New("TMP")
	start := time.Now()
	for err != nil {
		conn, err = grpc.Dial(prereqEndpoint, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			log.Panic("ERROR: Unable to connect to Prereq Endpoint: ", err)
			continue
		}
		server := NewPrereqValidationClient(conn)
		ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
		preResult, err = server.CheckPrereqs(ctx, &prereq)
		if err != nil {
			log.Panic("ERROR: Unable to connect to Prereq Endpoint: ", err)
		} else if time.Now().Sub(start).Milliseconds() > 1000 {
			return nil, errors.New("Timeout")
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

	for _, index := range preResult.InvalidIndices {
		dprint("Invalid Index: ", index)
		resultMap[index] = cvInterface.ResultClass_INVALID
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
