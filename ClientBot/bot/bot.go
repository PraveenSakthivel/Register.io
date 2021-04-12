package bot

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	cvInterface "registerio/clientBot/protobuf/cvInterface"
	db "registerio/clientBot/protobuf/dbRequests"
	rvInterface "registerio/clientBot/protobuf/rvInterface"
	login "registerio/clientBot/protobuf/token"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var sec credentials.TransportCredentials

//DB Info
const (
	host     = "prod-1-cluster.cluster-csz2smpfztf7.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

func init() {
	certPool, _ := x509.SystemCertPool()
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            certPool,
	}
	sec = credentials.NewTLS(config)
}

func loginRoutine(netid string, password string) (string, error) {
	//Hit Website
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://www.registerio.co")
	if err != nil {
		return "", err
	}
	//Login, get current registration
	conn, err := grpc.Dial("login.registerio.co:8080", grpc.WithTransportCredentials(sec))
	if err != nil {
		return "", err
	}

	defer conn.Close()
	server := login.NewLoginEndpointClient(conn)
	creds := login.Credentials{NetID: netid, Password: password}
	resp, err := server.GetLoginToken(context.Background(), &creds)
	token := resp.Token
	if err != nil {
		return "", err
	}
	tokenStruct := login.Token{Token: token}
	_, err = server.GetCurrentRegistrations(context.Background(), &tokenStruct)
	if err != nil {
		return "", err
	}
	return token, nil

}

func rvRoutine(token string) error {
	conn, err := grpc.Dial("rv.registerio.co:8080", grpc.WithTransportCredentials(sec))
	if err != nil {
		return err
	}

	defer conn.Close()
	server := rvInterface.NewRegistrationValidationClient(conn)

	student := rvInterface.Student{Token: token}

	_, err = server.CheckRegVal(context.Background(), &student)
	return err

}

func getClasses() ([]string, error) {
	var classes []string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT index FROM \"sqs queues\" ORDER BY random() limit 5")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var index string
		err = rows.Scan(&index)
		if err != nil {
			return nil, err
		}
		classes = append(classes, index)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return classes, err
}

func cvRoutine(token string, ops []*cvInterface.ClassOperations) error {
	conn, err := grpc.Dial("cv.registerio.co:8080", grpc.WithTransportCredentials(sec))
	if err != nil {
		log.Println("ERROR: Unable to connect")
		return err
	}

	defer conn.Close()
	server := cvInterface.NewCourseValidationClient(conn)
	req := cvInterface.RegistrationRequest{Token: token, Classes: ops}
	_, err = server.ChangeRegistration(context.Background(), &req)
	return err
}

func dbRoutine(token string, classes []string) error {
	req := db.ClassAddStatusParams{Token: token, Index: classes}
	conn, err := grpc.Dial("database.registerio.co:8080", grpc.WithTransportCredentials(sec))
	if err != nil {
		return err
	}
	defer conn.Close()
	server := db.NewDatabaseWrapperClient(conn)
	_, err = server.ClassAddStatus(context.Background(), &req)

	return err

}

func RunBot(netid string, password string, out chan error) {
	token, err := loginRoutine(netid, password)
	if err != nil {
		log.Println("Login")
		out <- err
		return
	}
	err = rvRoutine(token)
	if err != nil {
		log.Println("RV")
		out <- err
		return
	}

	classes, err := getClasses()
	if err != nil {
		out <- err
		return
	}

	var addOps []*cvInterface.ClassOperations
	var dropOps []*cvInterface.ClassOperations

	for _, class := range classes {
		addOps = append(addOps, &cvInterface.ClassOperations{Index: class, Op: cvInterface.ReqOp_ADD})
		dropOps = append(dropOps, &cvInterface.ClassOperations{Index: class, Op: cvInterface.ReqOp_DROP})
	}

	max := 42
	min := 20
	for i := 1; i <= 5; i++ {
		err = cvRoutine(token, addOps)
		if err != nil {
			log.Println("CV-ADD: ", err)
			out <- err
			return
		}
		time.Sleep(time.Second)
		err = dbRoutine(token, classes)
		if err != nil {
			log.Println("DB")
			out <- err
			return
		}
		time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Second)
		err = cvRoutine(token, dropOps)
		if err != nil {
			log.Println("DROP")
			out <- err
			return
		}
		time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Second)
	}
	err = cvRoutine(token, addOps)
	if err != nil {
		log.Println("CV-ADD")
		out <- err
		return
	}
	out <- nil
	return
}
