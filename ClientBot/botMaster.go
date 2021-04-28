package main

import (
	"log"
	"os"
	bot "registerio/clientBot/bot"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	env := os.Getenv("START")
	start, _ := strconv.Atoi(env)
	num := 200
	c := make(chan error)
	for i := start; i < start+num; i++ {
		go bot.RunBot("bot"+strconv.Itoa(i), "Password", c)
	}
	namespace := "Prod/BotErrors"
	name := "BotErrors"
	val := 1.0
	data := cloudwatch.MetricDatum{}
	data.MetricName = &name
	data.Value = &val
	metric := []*cloudwatch.MetricDatum{&data}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := cloudwatch.New(sess)
	for i := 1; i <= num; i++ {
		err := <-c
		if err != nil {
			log.Println("Error with bot", err)
			req, _ := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
				MetricData: metric,
				Namespace:  &namespace,
			})
			err := req.Send()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
