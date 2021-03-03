package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Prereq struct {
	Class string
	Grade int
}

const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

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

func GetPrereqs() (map[string][][]Prereq, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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

func GetLookups() (map[string]string, error) {
	retval := make(map[string]string)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
