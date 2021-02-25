package models

import (
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lib/pq"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User ...
type User struct {
	Netid    string
	Password string
	Type     int
}

// CourseRegistration ...
type CourseRegistration struct {
	Netid      string
	ClassIndex string
}

// Soc ...
type Soc struct {
	Location        string
	Level           string
	School          int
	Department      int
	ClassNumber     int
	Index           string
	Name            string
	Section         string
	MeetingLocation string
	MeetingTimes    string
	Exam            string
	Instructors     pq.StringArray `gorm:"type:character varying[]"`
	Codes           pq.StringArray `gorm:"type:character varying[]"`
	Synopsis        string
	Books           pq.StringArray `gorm:"type:character varying[]"`
}

// CourseHistory ...
type CourseHistory struct {
	Netid        string
	Semester     string
	CourseNumber string
	Grade        string
}

// Token ...
// type Token struct {
// 	gorm.Model
// 	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// 	Value   string
// 	Expires time.Time
// }

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

	database.AutoMigrate(&User{})
	database.AutoMigrate(&CourseRegistration{})
	database.AutoMigrate(&Soc{})
	database.AutoMigrate(&CourseHistory{})
	// database.Model(&Token{}).AddForeignKey("User", "users(Email)", "CASCADE", "CASCADE")
	// database.AutoMigrate(&Token{})
	DB = database

}
