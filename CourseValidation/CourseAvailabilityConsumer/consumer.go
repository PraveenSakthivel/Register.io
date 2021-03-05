package main

import (
	"flag"
	"log"
	data "registerio/cv/consumer/database"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var index = "02345"
var queueURL *string
var state data.Consumer
var svc *sqs.SQS
var timeout int64
var waitTime int64
var debug bool

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func setup() error {
	var err error
	state, err = data.RetrieveState(index)
	if err != nil {
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	if err != nil {
		log.Println("Error Generating AWS Session: ", err)
		return err
	}
	svc = sqs.New(sess)
	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(index + ".fifo"),
	})
	if err != nil {
		log.Println("Error Getting queue url: ", err)
		return err
	}

	queueURL = urlResult.QueueUrl
	timeout = 5
	waitTime = 20
	return nil
}

func retrieveMessages() ([]*sqs.Message, error) {
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

	if err != nil {
		log.Println("Error Pulling Messages from SQS: ", err)
		return nil, err
	}

	return (*msgResult).Messages, nil
}

func addStudent(netID string, spn bool) bool {
	if state.CurrentSize == state.MaxSize && !spn {
		dprint("ADD|Class full cannot add: ", netID)
		return false
	}

	for _, student := range state.RegisteredStudents {
		if student == netID {
			dprint("ADD|Student already in class: ", netID)
			return false
		}
	}

	state.RegisteredStudents = append(state.RegisteredStudents, netID)
	state.CurrentSize++
	dprint("ADD|Added Student: ", netID)
	return true
}

func dropStudent(netID string) bool {
	for i, student := range state.RegisteredStudents {
		if student == netID {
			students := &state.RegisteredStudents
			//Rewrite student with last element and shorten slice
			(*students)[i] = (*students)[len(*students)-1]
			state.RegisteredStudents = (*students)[:len(*students)-1]
			state.CurrentSize--
			dprint("DROP|Dropped Student: ", netID)
			return true
		}
	}
	dprint("DROP|Student not in class: ", netID)
	return false
}

func proccessMessage(message *sqs.Message) error {
	fields := strings.Split(*message.Body, "|")
	netID, action := fields[0], fields[1]
	var toUpdate bool
	switch action {
	case "add":
		toUpdate = addStudent(netID, false)
	case "drop":
		toUpdate = dropStudent(netID)
	case "spn":
		toUpdate = addStudent(netID, true)
	default:
		log.Println("Error Unknown Action: ", action)
		toUpdate = false
	}

	if toUpdate {
		err := data.UpdateState(state)
		if err != nil {
			log.Println("Error Updating Database: ", err)
			return err
		}
	}

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
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
	err := setup()
	if err != nil {
		return
	}

	for true {
		messages, err := retrieveMessages()
		if err != nil {
			continue
		}

		for _, message := range messages {
			for proccessMessage(message) != nil {
			}
		}

	}

}
