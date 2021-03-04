package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// JWTService ...
type JWTService interface {
	GenerateToken(netid string, isUser bool, classes map[string]int, cases []int) string
	ValidateToken(token string) (*jwt.Token, error)
}

// Custom fields we can expand on
type authCustomClaims struct {
	Name         string         `json:"name"`
	User         bool           `json:"user"`
	Classes      map[string]int `json:"classes"`
	SpecialCases []int          `json:"special"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

// JWTAuthService ..
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		// This will be changed
		issure: "Max",
	}
}

// Get secret from .env file
func getSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// generate token and seed with netid information
func (service *jwtServices) GenerateToken(netid string, isUser bool, classes map[string]int, cases []int) string {
	claims := &authCustomClaims{
		netid,
		isUser,
		classes,
		cases,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// parse token and ensure it is valid
func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
