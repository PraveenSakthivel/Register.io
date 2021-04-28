package main

import (
	"context"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
	"main/models"
	Tokens "main/protobuf"
	"net"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

var debug = false

// Server ...
type Server struct {
	Tokens.UnimplementedLoginEndpointServer
	debug bool
}

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

// NewServer -- Pulls Info from DB and creates new Server Struct
func NewServer() *Server {
	s := &Server{}
	models.ConnectDB()
	return s
}

// AddToList ...
func (s *Server) AddToList(ctx context.Context, in *Tokens.Addition) (*Tokens.Token, error) {
	netid := in.NetID
	course := in.Course
	entry := models.Watchlist{Netid: netid, Course: course}
	result := models.DB.Create(&entry)
	if result.Error != nil {
		return &Tokens.Token{}, errors.New("Error adding to watchlist")
	}
	return &Tokens.Token{Token: "Added to watchlist"}, nil
}

// DropFromList ...
func (s *Server) DropFromList(ctx context.Context, in *Tokens.Addition) (*Tokens.Token, error) {
	netid := in.NetID
	course := in.Course
	result := models.DB.Where("netid = ? AND course = ?", netid, course).Delete(&models.Watchlist{})
	if result.Error != nil {
		return &Tokens.Token{}, errors.New("Error removing from watchlist")
	}
	return &Tokens.Token{Token: "Removed from watchlist"}, nil
}

// PingList ...
func (s *Server) PingList(ctx context.Context, in *Tokens.Token) (*Tokens.Token, error) {
	course := in.Token
	smtpServer := "smtp.gmail.com"
	godotenv.Load(".env")
	auth := smtp.PlainAuth(
		"",
		"registeriowatchlist",
		os.Getenv("EMAILPASS"),
		"smtp.gmail.com",
	)
	entries := []models.Watchlist{}
	from := mail.Address{Name: "Register.io", Address: "registeriowatchlist@gmail.com"}
	models.DB.Where("course = ?", course).Find(&entries)

	for _, entry := range entries {
		recip := entry.Netid + "@scarletmail.rutgers.edu"
		to := mail.Address{Name: "", Address: recip}
		title := "Register.io=?utf-8?Q?=F0=9F=93=9A?= | Opening for " + entry.Course
		body := "There is currently an open space for the course: " + entry.Course + "\r\nThanks for using Register.io📚"
		header := make(map[string]string)
		header["From"] = from.String()
		header["To"] = to.String()
		header["Subject"] = title
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"
		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

		err := smtp.SendMail(
			smtpServer+":587",
			auth,
			from.Address,
			[]string{to.Address},
			[]byte(message),
		)
		if err != nil {
			log.Println(err)
			return &Tokens.Token{}, errors.New("Error pinging list list")
		}
	}

	return &Tokens.Token{Token: "Watchlist pinged!"}, nil
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
	Tokens.RegisterLoginEndpointServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
