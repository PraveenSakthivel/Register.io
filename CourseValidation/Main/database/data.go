package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	classTiming "registerio/cv/main/classtiming"
	secret "registerio/cv/main/secrets"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type SPN struct {
	User  string
	Index string
}

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

func (s *DB) GetQueues() (map[string]string, error) {
	retval := make(map[string]string)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}

	rows, err := db.Query("SELECT index, url FROM \"sqs queues\"")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var queue string
		var url string
		err = rows.Scan(&queue, &url)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		retval[queue] = url
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

func (s *DB) GetSPNs() (map[string]SPN, error) {
	retval := make(map[string]SPN)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}

	rows, err := db.Query("SELECT spn, \"class index\", \"user\" FROM spns")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var spn string
		var index string
		var user string
		err = rows.Scan(&spn, &index, &user)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		newSPN := SPN{User: user, Index: index}
		retval[spn] = newSPN
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

func (s *DB) GetClassTimes() (map[string][]*classTiming.ClassSlot, error) {
	retval := make(map[string][]*classTiming.ClassSlot)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}

	rows, err := db.Query("SELECT location, index,\"meeting times\" FROM soc;")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var location string
		var index string
		var times string
		err = rows.Scan(&location, &index, &times)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		slots, err := classTiming.BuildClassSlots(times, location)
		if err != nil {
			log.Println("Error Building Class Slots: ", err)
			return nil, err
		}
		retval[index] = slots
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

func (s *DB) GetCurrentRegistration(netID string) ([]string, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}

	query := `SELECT netid, ARRAY_AGG("class index")
	FROM "course registration" WHERE netid = $1 GROUP BY netid;`

	rows, err := db.Query(query, netID)
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user string
		var indices []string
		err = rows.Scan(&user, pq.Array(&indices))
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		return indices, nil
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return []string{}, nil
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
