package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	rvInterface "registerio/rv/protobuf"
	secret "registerio/rv/secrets"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

var debug = false

//DB Info
type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

type Server struct {
	rvInterface.UnimplementedRegistrationValidationServer
	students    map[string]int
	debug       bool
	tokenSecret string
	db          *DB
}

type userClaims struct {
	NetID        string           `json:"name"`
	ClassHistory map[string]int32 `json:"classHistory"`
	SpecialCases map[string]bool  `json:"specialCases"`
	jwt.StandardClaims
}

type token struct {
	TokenSecret string
}

func dprint(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func buildDB() (*DB, error) {
	dbstring, err := secret.GetTokenSecret("prod/DB")
	if err != nil {
		return nil, err
	}
	retval := DB{}
	err = json.Unmarshal([]byte(dbstring), &retval)
	if err != nil {
		return nil, err
	}
	return &retval, nil
}

func (s *Server) Decrypt(encryptedString string) (string, error) {
	// key, _ := hex.DecodeString(tokenObj.Token[0:32])
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher([]byte(s.tokenSecret[0:32]))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	if len(enc) < nonceSize {
		return "", err
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}

//Add secret decoding and check for validity
func (s *Server) parseJWT(encodedToken string) (string, error) {
	decodedToken, err := s.Decrypt(encodedToken)
	dprint(decodedToken)
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(decodedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(s.tokenSecret), nil
	})

	if err != nil {
		log.Println("Error Parsing JWT: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(*userClaims); ok {
		specCases := make(map[int32]bool)
		for key, val := range claims.SpecialCases {
			intkey, err := strconv.ParseInt(key, 10, 32)
			if err != nil {
				log.Println("Error Parsing Cases: ", err)
				continue
			}
			specCases[int32(intkey)] = val
		}
		dprint(claims.NetID)
		return claims.NetID, nil
	}
	return "", errors.New("Unable to Parse JWT")
}

// Retrieve list of all students from Database
// TODO: Retrieve endpoint securely
func (s *DB) retrieveData() map[string]int {
	students := make(map[string]int)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}

	rows, err := db.Query("SELECT * FROM \"registration dates\"")
	if err != nil {
		log.Println("Database error: ", err)
		os.Exit(3)
	}
	defer rows.Close()
	for rows.Next() {
		var netid string
		var time int
		err = rows.Scan(&netid, &time)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			os.Exit(3)
		}
		students[netid] = time
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		os.Exit(3)
	}

	dprint("OK: Successfully Pulled Data from DB")
	return students
}

//Pulls Info from DB and creates new Server Struct
func NewServer() *Server {
	db, err := buildDB()
	if err != nil {
		log.Fatal("Error building database: ", err)
	}
	students := db.retrieveData()
	tokenSecret, err := secret.GetTokenSecret("user/JWTEncryption")
	if err != nil {
		log.Fatal("ERROR: Cannot get token secret: ", err)
	}
	var Token token
	json.Unmarshal([]byte(tokenSecret), &Token)
	s := &Server{students: students, debug: debug, db: db, tokenSecret: Token.TokenSecret}
	return s
}

//Checks whether user is elgibile to register
func (s *Server) CheckRegVal(ctx context.Context, student *rvInterface.Student) (*rvInterface.Response, error) {
	resp := rvInterface.Response{
		Eligible: false,
		Time:     -1,
	}

	//Parse token
	netID, err := s.parseJWT(student.Token)
	if err != nil {
		log.Println("ERROR: Invalid Token")
		return &resp, err
	}
	// Check to see if student is eligible
	if dateInt, ok := s.students[netID]; ok {
		resp.Time = int64(dateInt)
		date := time.Unix(resp.Time, 0)
		if time.Now().After(date) {
			resp.Eligible = true
		}
	} else {
		log.Println("WARNING: Unidentifiable NetID ", netID)
	}

	dprint("OK: Request with NetID: ", netID)
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
	rvInterface.RegisterRegistrationValidationServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to listen on port 8080: ", err)
	}
}
