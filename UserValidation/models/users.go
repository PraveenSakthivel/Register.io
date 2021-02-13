package models

import (
	"time"

	"github.com/jinzhu/gorm"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User ...
type User struct {
	Email    string
	Password string
}

// Token ...
type Token struct {
	gorm.Model
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Value   string
	Expires time.Time
}

// DB ...
var DB *gorm.DB

// ConnectDB ...
func ConnectDB() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to DB")
	}

	database.AutoMigrate(&User{})
	// database.Model(&Token{}).AddForeignKey("User", "users(Email)", "CASCADE", "CASCADE")
	database.AutoMigrate(&Token{})
	DB = database

}
