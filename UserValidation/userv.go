package main

import (
	"context"
	"flag"
	"log"
	"main/controller"
	"main/models"
	Tokens "main/protobuf"
	"main/service"
	"net"

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

// GetLoginToken -- logs user in and returns token
func (s *Server) GetLoginToken(ctx context.Context, creds *Tokens.Credentials) (*Tokens.Token, error) {
	resp := Tokens.Token{Token: ""}
	token := s.LoginController.LoginEndpoint(creds)
	if token != "" {
		resp.Token = token
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
