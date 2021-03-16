package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	classTiming "registerio/cv/main/classtiming"
	data "registerio/cv/main/database"
	prereqInterface "registerio/cv/main/prereq"
	cvInterface "registerio/cv/main/protobuf"
	secret "registerio/cv/main/secrets"

	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

var debug = false

//Server instance
type Server struct {
	cvInterface.UnimplementedCourseValidationServer
	svc         *sqs.SQS
	queueLookup map[string]string
	spns        map[string]data.SPN
	timings     map[string][]*classTiming.ClassSlot
	tokenSecret string
	db          *data.DB
}

//Student representation for handling request
type Student struct {
	netID        string
	classHistory map[string]int32
	cases        map[int32]bool
}

type classResult struct {
	index string
	err   error
}

type userClaims struct {
	NetID        string           `json:"name"`
	ClassHistory map[string]int32 `json:"classHistory"`
	SpecialCases map[string]bool  `json:"specialCases"`
	jwt.StandardClaims
}

type token struct {
	TokenSecret string
}

//debug print
func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

//Initialize any Server Variables
func NewServer() *Server {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	db, err := data.BuildDB()
	if err != nil {
		log.Fatal("ERROR: Cannot retrieve Database secret")
	}
	queues, err := db.GetQueues()
	if err != nil {
		log.Fatal("ERROR: Cannot retrieve queue lookup table")
	}
	spns, err := db.GetSPNs()
	if err != nil {
		log.Fatal("ERROR: Cannot retrieve SPNs")
	}
	timings, err := db.GetClassTimes()
	if err != nil {
		log.Fatal("ERROR: Cannot get class times")
	}
	tokenSecret, err := secret.GetTokenSecret("user/JWTEncryption")
	if err != nil {
		log.Fatal("ERROR: Cannot get token secret: ", err)
	}
	var Token token
	json.Unmarshal([]byte(tokenSecret), &Token)

	s := &Server{svc: svc, queueLookup: queues, spns: spns, timings: timings, tokenSecret: Token.TokenSecret, db: db}
	return s
}

//Add secret decoding and check for validity
func (s *Server) parseJWT(encodedToken string) (Student, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(s.tokenSecret), nil
	})

	if err != nil {
		log.Println("Error Parsing JWT: ", err)
		return Student{}, err
	}

	if claims, ok := token.Claims.(*userClaims); ok {
		specCases := make(map[int32]bool)
		for key, val := range claims.SpecialCases {
			intkey, err := strconv.ParseInt(key, 10, 32)
			if err != nil {
				log.Println("Error Parsing Cases: ", err)
				continue
			}
			specCases[int32(intkey)] = val
		}
		return Student{netID: claims.NetID, classHistory: claims.ClassHistory, cases: specCases}, nil
	}
	return Student{}, errors.New("Unable to Parse JWT")
}

func (s *Server) sendRegRequest(netID string, class *cvInterface.ClassOperations, c chan classResult) {
	if _, ok := s.queueLookup[class.Index]; !ok {
		c <- classResult{index: class.Index, err: errors.New("Unable to find SQS Queue url")}
		return
	}
	url := s.queueLookup[class.Index]
	var opString string
	switch class.Op {
	case cvInterface.ReqOp_ADD:
		opString = "add"
	case cvInterface.ReqOp_DROP:
		opString = "drop"
	case cvInterface.ReqOp_SPN:
		opString = "spn"
	default:
		c <- classResult{index: class.Index, err: errors.New("Unsupported Operation")}
		return
	}
	message := netID + "|" + opString
	start := time.Now()
	dedupID := netID + class.Index + strconv.Itoa(int(start.Unix()))
	dprint(dedupID)
	for true {
		_, err := s.svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds:           aws.Int64(0),
			MessageBody:            &message,
			QueueUrl:               &url,
			MessageGroupId:         aws.String(class.Index),
			MessageDeduplicationId: &dedupID,
		})
		if err == nil {
			break
		} else if time.Now().Sub(start).Milliseconds() > 1000 {
			c <- classResult{index: class.Index, err: errors.New("Timeout")}
			return
		}
		log.Println("Error Sending SQS Message: ", err)
	}
	c <- classResult{index: class.Index, err: nil}
	return

}

func (s *Server) getCurrentSchedule(netID string) (map[time.Weekday]*classTiming.ClassSlot, error) {
	currentRegistration, err := s.db.GetCurrentRegistration(netID)
	if err != nil {
		return nil, err
	}
	return classTiming.BuildSchedule(currentRegistration, &s.timings)
}

func (s *Server) checkSchedandSend(operation *cvInterface.ClassOperations, currentSched map[time.Weekday]*classTiming.ClassSlot, netID string, c *(chan classResult)) (bool, error) {
	classTime := s.timings[operation.Index]
	good, err := classTiming.CheckTimesAndInsert(classTime, currentSched)
	if err != nil {
		return false, err
	} else if !good {
		return false, nil
	}
	go s.sendRegRequest(netID, operation, *c)
	return true, nil
}

func (s *Server) AddSPN(ctx context.Context, req *cvInterface.SPNRequest) (*cvInterface.SPNResponse, error) {
	response := cvInterface.SPNResponse{Valid: false, Result: cvInterface.ResultClass_ERROR}
	student, err := s.parseJWT(req.Token)
	dprint("Received SPN Request for User: ", student)
	if err != nil {
		log.Panic("ERROR Parsing JWT")
		return &response, err
	}

	if spn, ok := s.spns[req.Spn]; ok {
		if spn.User == student.netID && spn.Index == req.Index {
			dprint("SPN Match")
			response.Valid = true
			currentSched, err := s.getCurrentSchedule(student.netID)
			if err != nil {
				return &response, err
			}
			c := make(chan classResult)
			op := cvInterface.ClassOperations{Index: spn.Index, Op: cvInterface.ReqOp_SPN}
			eligible, err := s.checkSchedandSend(&op, currentSched, student.netID, &c)
			if err != nil {
				return &response, err
			} else if err == nil && !eligible {
				response.Result = cvInterface.ResultClass_TIME
				return &response, nil
			}
			result := <-c
			if result.err != nil {
				response.Result = cvInterface.ResultClass_OK
				dprint("Sent Message to SQS: ", result.index)
			}
			return &response, result.err
		}
	}

	return &response, nil
}

func (s *Server) ChangeRegistration(ctx context.Context, req *cvInterface.RegistrationRequest) (*cvInterface.RegistrationResponse, error) {
	results := make(map[string]cvInterface.ResultClass)
	var indices []string
	var response cvInterface.RegistrationResponse
	response.Results = make(map[string]cvInterface.ResultClass)
	classes := req.Classes
	for _, class := range classes {
		response.Results[class.Index] = cvInterface.ResultClass_ERROR
		indices = append(indices, class.Index)
	}

	student, err := s.parseJWT(req.Token)
	dprint("Received Request for User: ", student)
	if err != nil {
		return &response, err
	}

	currentSched, err := s.getCurrentSchedule(student.netID)
	if err != nil {
		return &response, err
	}

	preResult, err := prereqInterface.CheckPrereqs(student.classHistory, indices, student.cases)
	if err != nil {
		log.Panic("ERROR Timed out trying to connect to prereq endpoint")
		return &response, err
	}
	dprint("Checking Prereqs")
	eligibleClasses := prereqInterface.EvalPrereqResults(preResult, &results, req.Classes, debug)
	c := make(chan classResult)

	classReqsSent := 0

	for _, eligibleClass := range eligibleClasses {
		eligible, err := s.checkSchedandSend(eligibleClass, currentSched, student.netID, &c)
		//Request response defaults to error so no need to update
		if err == nil && !eligible {
			response.Results[eligibleClass.Index] = cvInterface.ResultClass_TIME
		} else if err == nil && eligible {
			classReqsSent++
		}
	}

	for i := 0; i < classReqsSent; i++ {
		result := <-c
		if result.err != nil {
			results[result.index] = cvInterface.ResultClass_ERROR
		} else {
			dprint("Sent Message to SQS: ", result.index)
			results[result.index] = cvInterface.ResultClass_OK
		}
	}

	response = cvInterface.RegistrationResponse{Results: results}
	return &response, nil

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
