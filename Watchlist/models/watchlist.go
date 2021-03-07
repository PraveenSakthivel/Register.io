package models

import (
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Watchlist ...
type Watchlist struct {
	Netid  string
	Course string
}

// DB ...
var DB *gorm.DB

// ConnectDB ...
func ConnectDB() {
	godotenv.Load(".env")
	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("DBUSER"), os.Getenv("DBPASS")),
		Scheme:   "postgres",
		Host:     "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com:5432",
		Path:     "maindb",
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
