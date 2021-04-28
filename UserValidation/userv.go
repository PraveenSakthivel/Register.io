package main

import (
	"context"
	"errors"
	"flag"
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
		user := models.User{}
		models.DB.Where("netid = ?", token.Claims.(jwt.MapClaims)["name"]).First(&user)
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
			reg.MeetingTimes = class.MeetingTimes
			reg.Exam = class.Exam
			reg.Instructors = class.Instructors
			reg.Codes = class.Codes
			reg.Synopsis = class.Synopsis
			reg.Books = class.Books
			reg.Credits = int64(class.Credits)
			result.Classes = append(result.Classes, &reg)

		}
		result.UserType = int64(user.Type)
		// fmt.Println(classes)

		return &result, nil
	}
	return &Tokens.Registrations{}, errors.New("Invalid user token")
}

// GetLoginToken -- logs user in and returns token
func (s *Server) GetLoginToken(ctx context.Context, creds *Tokens.Credentials) (*Tokens.Response, error) {
	resp := Tokens.Response{Token: "", UserType: -1}
	token := s.LoginController.LoginEndpoint(creds)
	if token != "" {
		user := models.User{}
		models.DB.Where("netid = ?", creds.NetID).First(&user)
		resp.Token = middleware.Encrypt(token)
		resp.UserType = int64(user.Type)
		// data, _ := proto.Marshal(&resp)
		// stringarray := fmt.Sprint(data)
		// stringarray = stringarray[1 : len(stringarray)-1]
		// return &Tokens.Response{Token: stringarray, UserType: int64(user.Type)}, nil
		return &resp, nil
	}
	log.Println("WARNING: Invalid credentials", creds.NetID)
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
