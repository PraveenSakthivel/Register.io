package models

import (
	"encoding/json"
	secret "main/secrets"
	"net/url"

	"github.com/jinzhu/gorm"
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

// Students ...
type Students struct {
	Netid        string
	Class        int
	Credits      int
	SpecialCases pq.Int64Array `gorm:"type:integer[]"`
}

// CourseRegistration ...
type CourseRegistration struct {
	Netid      string
	ClassIndex string
}

// Tabler ...
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Soc) TableName() string {
	return "soc"
}

// Soc ...
type Soc struct {
	Location        string
	Level           string
	School          int
	Department      int
	ClassNumber     int `gorm:"column:class number"`
	Index           string
	Name            string
	Section         string
	MeetingLocation string `gorm:"column:meeting location"`
	MeetingTimes    string `gorm:"column:meeting times"`
	Exam            string
	Instructors     pq.StringArray `gorm:"type:character varying[]"`
	Codes           pq.StringArray `gorm:"type:character varying[]"`
	Synopsis        string
	Books           pq.StringArray `gorm:"type:character varying[]"`
	Credits         int
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

	database.AutoMigrate(&User{})
	database.AutoMigrate(&CourseRegistration{})
	database.AutoMigrate(&Soc{})
	database.AutoMigrate(&CourseHistory{})
	database.AutoMigrate(&Students{})
	// database.Model(&Token{}).AddForeignKey("User", "users(Email)", "CASCADE", "CASCADE")
	// database.AutoMigrate(&Token{})
	DB = database

}
