package service

import (
	"encoding/json"
	"fmt"
	"time"

	secret "main/secrets"

	"github.com/dgrijalva/jwt-go"
)

// JWTService ...
type JWTService interface {
	GenerateToken(netid string, isUser bool, classes map[string]int, cases map[int]bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

// Custom fields we can expand on
type authCustomClaims struct {
	Name         string         `json:"name"`
	User         bool           `json:"user"`
	ClassHistory map[string]int `json:"classhistory"`
	SpecialCases map[int]bool   `json:"special"`
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

type TokenSecret struct {
	Token string `json:"TokenSecret"`
}

// Get secret from .env file
func getSecretKey() string {
	tokenSecret, _ := secret.GetTokenSecret("user/JWTEncryption")
	tokenObj := TokenSecret{}
	json.Unmarshal([]byte(tokenSecret), &tokenObj)
	if tokenObj.Token == "" {
		return "secret"
	}
	return tokenObj.Token
}

// generate token and seed with netid information
func (service *jwtServices) GenerateToken(netid string, isUser bool, classes map[string]int, cases map[int]bool) string {
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
