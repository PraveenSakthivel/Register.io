package main

import (
	"errors"
	"flag"
	"log"
	"os"
	data "registerio/cv/consumer/database"
	"strings"

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

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func setup() {
	var err error
	index = os.Getenv("INDEX")
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
	waitTime = 10
	return
}

func retrieveMessages() ([]*sqs.Message, error) {
	dprint("start")
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &waitTime,
	})
	dprint("here")

	if err != nil {
		log.Println("Error Pulling Messages from SQS: ", err)
		return nil, err
	}

	return (*msgResult).Messages, nil
}

func addStudent(netID string, spn bool) error {
	if state.CurrentSize == state.MaxSize && !spn {
		dprint("ADD|Class full cannot add: ", netID)
		return nil
	}

	for _, student := range state.RegisteredStudents {
		if student == netID {
			dprint("ADD|Student already in class: ", netID)
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
	dprint("ADD|Added Student: ", netID)
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
			dprint("DROP|Dropped Student: ", netID)
			return nil
		}
	}
	dprint("DROP|Student not in class: ", netID)
	return nil
}

func proccessMessage(message *sqs.Message) error {
	dprint("Received Message: ", *message.Body)
	fields := strings.Split(*message.Body, "|")
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
		ReceiptHandle: message.ReceiptHandle,
	})
	if err != nil {
		log.Println("Error Deleting Message from SQS: ", err)
	}

	return err
}

func main() {
	debugPrnt := flag.Bool("debug", false, "Debug Print all Requests")
	flag.Parse()
	debug = *debugPrnt
	setup()

	dprint(*queueURL)

	for true {
		messages, err := retrieveMessages()
		if err != nil {
			log.Println("Error retrieving messages: ", err)
			continue
		}
		for _, message := range messages {
			for proccessMessage(message) != nil {
			}
		}

	}

}
