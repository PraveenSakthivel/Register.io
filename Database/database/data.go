package data

import (
	"encoding/json"
	secret "registerio/db/secrets"
)

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

type token struct {
	TokenSecret string
}


func BuildDB() (*DB, error) {
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

func GetTokenSecret() (token, error) {
	tokenSecret, err := secret.GetTokenSecret("user/JWTEncryption")
	if err != nil {
		return token{}, err
	}
	var Token token
	json.Unmarshal([]byte(tokenSecret), &Token)
	return Token, nil
}
