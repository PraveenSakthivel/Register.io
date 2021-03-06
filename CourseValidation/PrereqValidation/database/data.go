package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	secret "registerio/cv/preqreq/secrets"

	"github.com/lib/pq"
)

type Prereq struct {
	Class string
	Grade int
}

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

func processPrereqs(rows *sql.Rows) (map[string][][]Prereq, error) {
	retval := make(map[string][][]Prereq)
	var master [][]Prereq //List of all prereq sets
	var currentCourse string
	for rows.Next() {
		var course string
		var reqs []string
		var grades []float32
		err := rows.Scan(&course, pq.Array(&reqs), pq.Array(&grades))
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		var preqreqs []Prereq
		//Create a set of prereqs
		for i := range reqs {
			new := Prereq{Class: reqs[i], Grade: int(grades[i] * 100)}
			preqreqs = append(preqreqs, new)
		}
		if currentCourse == "" {
			currentCourse = course
		} else if currentCourse != course { //Finished parsing current course, save value and start new list
			retval[currentCourse] = master
			currentCourse = course
			master = nil
		}
		master = append(master, preqreqs)
	}
	retval[currentCourse] = master

	err := rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

func (s *DB) GetPrereqs() (map[string][][]Prereq, error) {

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

	rows, err := db.Query("SELECT \"course number\", prereq, grade FROM prereqs ORDER BY \"course number\"")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	return processPrereqs(rows)
}

func (s *DB) GetLookups() (map[string]string, error) {
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

	rows, err := db.Query("SELECT index, class FROM \"index class lookup\"")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var index string
		var course string
		err = rows.Scan(&index, &course)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		retval[index] = course
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
}

func (s *DB) GetSpecialCases() (map[string][]int32, error) {
	retval := make(map[string][]int32)

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

	rows, err := db.Query("SELECT class, cases FROM \"class special cases\"")
	if err != nil {
		log.Println("Database error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course string
		var cases pq.Int32Array
		err = rows.Scan(&course, &cases)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return nil, err
		}
		retval[course] = cases
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return nil, err
	}

	return retval, nil
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
