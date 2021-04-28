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

// DBStruct ...
type DBStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

// BuildingLookup ...
type BuildingLookup struct {
	BldgNum string
	Code    string
	Lat     string
	Long    string
	Name    string
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

// CourseRegistration ...
type CourseRegistration struct {
	Netid      string
	ClassIndex string
}

// BuildDB ...
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

	database.AutoMigrate(&BuildingLookup{}, &CourseRegistration{}, &Soc{})
	DB = database

}
