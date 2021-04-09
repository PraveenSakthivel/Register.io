package bot

import (
	"context"
	"credentials"
	"crypto/x509"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	cvInterface "registerio/clientBot/protobuf/cvInterface"
	rvInterface "registerio/clientBot/protobuf/rvInterface"
	login "registerio/clientBot/protobuf/token"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var tls credentials.TransportCredentials

//DB Info
const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
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
	tls = credentials.NewTLS(config)
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
	conn, err := grpc.Dial("login.registerio.co:8080", grpc.WithTransportCredentials(tls))
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
	conn, err := grpc.Dial("rv.registerio.co:8080", grpc.WithTransportCredentials(tls))
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
	conn, err := grpc.Dial("cv.registerio.co:8080", grpc.WithTransportCredentials(tls))
	if err != nil {
		return err
	}

	defer conn.Close()
	server := cvInterface.NewCourseValidationClient(conn)
	req := cvInterface.RegistrationRequest{Token: token, Classes: ops}
	_, err = server.ChangeRegistration(context.Background(), &req)
	return err
}

func runBot(netid string, password string) error {
	token, err := loginRoutine(netid, password)
	if err != nil {
		return err
	}
	err = rvRoutine(token)
	if err != nil {
		return err
	}

	classes, err := getClasses()
	if err != nil {
		return err
	}

	var addOps []*cvInterface.ClassOperations
	var dropOps []*cvInterface.ClassOperations

	for _, class := range classes {
		addOps = append(addOps, &cvInterface.ClassOperations{Index: class, Op: cvInterface.ReqOp_ADD})
		dropOps = append(dropOps, &cvInterface.ClassOperations{Index: class, Op: cvInterface.ReqOp_DROP})
	}

	for i := 1; i <= 5; i++ {
		cvRoutine(token, addOps)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		cvRoutine(token, dropOps)
	}
	return nil
}
