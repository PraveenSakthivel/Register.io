package main

import (
	"context"
	"errors"
	"log"
	"os"
	data "registerio/cv/consumer-lambda/database"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var index string
var queueURL *string
var state data.Consumer
var svc *sqs.SQS
var timeout int64
var waitTime int64
var debug bool
var db *data.DB

func init() {
	var err error
	index = os.Args[1]
	log.Println("Building Database")
	db, err = data.BuildDB()
	if err != nil {
		log.Fatal("ERROR Unable to build DB: ", err)
	}

	state, err = db.RetrieveState(index)
	if err != nil {
		log.Fatal("ERROR Unable to retrieve state: ", err)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	if err != nil {
		log.Fatal("Error Generating AWS Session: ", err)
	}
	svc = sqs.New(sess)
	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(index + ".fifo"),
	})
	if err != nil {
		log.Fatal("Error Getting queue url: ", err)
	}

	queueURL = urlResult.QueueUrl
	timeout = 1
	waitTime = 1
}

func addStudent(netID string, spn bool) error {
	if state.CurrentSize == state.MaxSize && !spn {
		return nil
	}

	for _, student := range state.RegisteredStudents {
		if student == netID {
			return nil
		}
	}
	err := db.AddRegistration(netID, state.Index)
	if err != nil {
		log.Println("ERROR: Cannot add student: ", netID)
		return errors.New("Cannot add student")
	}

	state.RegisteredStudents = append(state.RegisteredStudents, netID)
	state.CurrentSize++
	return nil
}

func dropStudent(netID string) error {
	for i, student := range state.RegisteredStudents {
		if student == netID {
			err := db.RemoveRegistration(netID, state.Index)
			if err != nil {
				log.Println("ERROR: Cannot remove student: ", netID)
				return errors.New("Cannot remove student")
			}
			students := &state.RegisteredStudents
			//Rewrite student with last element and shorten slice
			(*students)[i] = (*students)[len(*students)-1]
			state.RegisteredStudents = (*students)[:len(*students)-1]
			state.CurrentSize--
			return nil
		}
	}
	return nil
}

func proccessMessage(body string, receiptHandle *string) error {
	fields := strings.Split(body, "|")
	netID, action := fields[0], fields[1]
	var err error
	switch action {
	case "add":
		err = addStudent(netID, false)
	case "drop":
		err = dropStudent(netID)
	case "spn":
		err = addStudent(netID, true)
	default:
		log.Println("Error Unknown Action: ", action)
	}

	if err != nil {
		return err
	}

	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		log.Println("Error Deleting Message from SQS: ", err)
	}

	return err
}

func HandleRequest(ctx context.Context, req events.SQSEvent) error {
	for _, message := range req.Records {
		err := proccessMessage(message.Body, &message.ReceiptHandle)
		if err != nil {
			return err
		}
	}
	db.PushState(state)
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
