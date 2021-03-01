package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	prereqInterface "registerio/cv/main/prereq"
	cvInterface "registerio/cv/main/protobuf"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var debug = false

const PrereqEndpoint = ":8081"

const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

//Server instance
type Server struct {
	cvInterface.UnimplementedCourseValidationServer
	svc         *sqs.SQS
	queueLookup map[string]string
}

//Student representation for handling request
type Student struct {
	netID        string
	classHistory map[string]int32
}

type classResult struct {
	index string
	err   error
}

type userClaims struct {
	NetID        string           `json:"name"`
	ClassHistory map[string]int32 `json:"classHistory"`
	jwt.StandardClaims
}

//debug print
func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func getQueues() (map[string]string, error) {
	retval := make(map[string]string)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}

	rows, err := db.Query("SELECT index, url FROM \"sqs queues\"")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var queue string
		var url string
		err = rows.Scan(&queue, &url)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		retval[queue] = url
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

//Initialize any Server Variables
func NewServer() *Server {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	queues, err := getQueues()
	if err != nil {
		log.Fatal("ERROR: Cannot retrieve queue lookup table")
	}
	s := &Server{svc: svc, queueLookup: queues}
	return s
}

//Add secret decoding and check for validity
func parseJWT(encodedToken string) (Student, error) {

	secret := "55a441b7b7fea3448945d090e0e67b79"

	token, err := jwt.ParseWithClaims(encodedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Error Parsing JWT: ", err)
		return Student{}, err
	}

	if claims, ok := token.Claims.(*userClaims); ok {
		return Student{netID: claims.NetID, classHistory: claims.ClassHistory}, nil
	}
	return Student{}, errors.New("Unable to Parse JWT")
}

func (s *Server) sendRegRequest(netID string, class *cvInterface.ClassOperations, c chan classResult) {
	url := s.queueLookup[class.Index]
	var opString string
	switch class.Op {
	case cvInterface.ReqOp_ADD:
		opString = "add"
	case cvInterface.ReqOp_DROP:
		opString = "drop"
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

func checkPrereqs(classHistory map[string]int32, indices []string) (*prereqInterface.PrereqResponse, error) {
	prereq := prereqInterface.PrereqRequest{ClassHistory: classHistory, Indices: indices}
	var preResult *prereqInterface.PrereqResponse
	var conn *grpc.ClientConn
	var err = errors.New("TMP")
	start := time.Now()
	for err != nil {
		conn, err = grpc.Dial(PrereqEndpoint, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			log.Panic("ERROR: Unable to connect to Prereq Endpoint: ", err)
			continue
		}
		server := prereqInterface.NewPrereqValidationClient(conn)
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

func evalPrereqResults(preResult *prereqInterface.PrereqResponse, results *map[string]cvInterface.ResultClass, classes []*cvInterface.ClassOperations) []*cvInterface.ClassOperations {
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

	student, err := parseJWT(req.Token)
	dprint("Received Request for User: ", student)
	if err != nil {
		log.Panic("ERROR Parsing JWT")
		return &response, nil
	}

	preResult, err := checkPrereqs(student.classHistory, indices)
	if err != nil {
		log.Panic("ERROR Timed out trying to connect to prereq endpoint")
		return &response, nil
	}
	dprint("Checking Prereqs")
	eligibleClasses := evalPrereqResults(preResult, &results, req.Classes)
	c := make(chan classResult)

	for _, eligibleClass := range eligibleClasses {
		go s.sendRegRequest(student.netID, eligibleClass, c)
	}

	for range eligibleClasses {
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
