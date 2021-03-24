package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"main/controller"
	"main/middleware"
	"main/models"
	Tokens "main/protobuf"
	"main/service"
	"net"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var debug = false

// Server ...
type Server struct {
	Tokens.UnimplementedLoginEndpointServer
	debug           bool
	JwtService      service.JWTService
	LoginController controller.LoginController
}

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

// NewServer -- Pulls Info from DB and creates new Server Struct
func NewServer() *Server {
	s := &Server{JwtService: service.JWTAuthService()}
	s.LoginController = controller.LoginHandler(s.JwtService)
	models.ConnectDB()
	return s
}

// GetCurrentRegistrations ...
func (s *Server) GetCurrentRegistrations(ctx context.Context, in *Tokens.Token) (*Tokens.Registrations, error) {
	token, valid := middleware.ValidTokenGRPC(in)

	if valid {

		regs := []models.CourseRegistration{}
		models.DB.Where("netid = ?", token.Claims.(jwt.MapClaims)["name"]).Find(&regs)
		// fmt.Println(regs)
		classes := []models.Soc{}
		for _, reg := range regs {
			current := []models.Soc{}
			models.DB.Where("index = ?", reg.ClassIndex).First(&current)
			classes = append(classes, current...)
		}
		result := Tokens.Registrations{}
		for _, class := range classes {
			reg := Tokens.Class{}
			reg.Location = class.Location
			reg.Level = class.Level
			reg.School = int64(class.School)
			reg.Department = int64(class.Department)
			reg.ClassNumber = int64(class.ClassNumber)
			reg.Index = class.Index
			reg.Name = class.Name
			reg.Section = class.Section
			reg.MeetingLocation = class.MeetingLocation
			reg.Exam = class.Exam
			reg.Instructors = class.Instructors
			reg.Codes = class.Codes
			reg.Synopsis = class.Synopsis
			reg.Books = class.Books
			result.Classes = append(result.Classes, &reg)

		}
		// fmt.Println(classes)

		return &result, nil
	}
	return &Tokens.Registrations{}, errors.New("Invalid user token")
}

// GetLoginToken -- logs user in and returns token
func (s *Server) GetLoginToken(ctx context.Context, creds *Tokens.Credentials) (*Tokens.Token, error) {
	resp := Tokens.Token{Token: ""}
	token := s.LoginController.LoginEndpoint(creds)
	if token != "" {

		resp.Token = middleware.Encrypt(token)
		data, _ := proto.Marshal(&resp)
		stringarray := fmt.Sprint(data)
		stringarray = stringarray[1 : len(stringarray)-1]
		return &Tokens.Token{Token: stringarray}, nil
	} else {
		log.Println("WARNING: Invalid credentials", creds.NetID)
	}

	dprint("OK: Logged in with NetID: ", creds.NetID)
	return &resp, nil
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
