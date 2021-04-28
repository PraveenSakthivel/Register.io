package models

import (
	"encoding/json"
	secret "main/secrets"
	"net/url"

	"github.com/jinzhu/gorm"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Watchlist ...
type Watchlist struct {
	Netid  string
	Course string
}

// DB ...
type DBStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

func BuildDB() (*DBStruct, error) {
	dbstring, err := secret.GetTokenSecret("prod/DB")
	if err != nil {
		return nil, err
	}
	retval := DBStruct{}
	err = json.Unmarshal([]byte(dbstring), &retval)
	if err != nil {
		return nil, err
	}
	return &retval, nil
}

// DB ...
var DB *gorm.DB

// ConnectDB ...
func ConnectDB() {

	dbobj, _ := BuildDB()

	dsn := url.URL{
		User:     url.UserPassword(dbobj.Username, dbobj.Password),
		Scheme:   "postgres",
		Host:     dbobj.Host + ":" + dbobj.Port,
		Path:     dbobj.Dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	database, err := gorm.Open("postgres", dsn.String())
	// database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to DB")
	}

	database.AutoMigrate(&Watchlist{})
	DB = database

}
