package models 

import (
	"database/sql"
	"log"
	"fmt"
	data "registerio/db/database"
)

func CheckStatus(s *data.DB, netID string, index string) (int, error) {
	
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return -1, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return -1, err
	}
	var query = "select exists(select 1 from course_registrations where netid=$1 and class_index=$2);"
	
	var exists bool
	err = db.QueryRow(query, netID, index).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
			log.Println("error checking if row exists for course registrations")
			return -1, err
    }
	if exists == true {
		return 1, nil
	}

	var sizeQuery = "SELECT ca.\"max size\" - COALESCE(crr.students, 0) FROM \"course availability\""+
	" ca INNER JOIN (SELECT cr.class_index, COUNT(cr.netid) as \"students\" FROM "+ "course_registrations cr WHERE cr.class_index = $1 GROUP BY cr.class_index) crr "+
	"ON ca.index = crr.class_index;"
	var spotsLeft int
	err = db.QueryRow(sizeQuery, index).Scan(&spotsLeft)
	if err != nil && err != sql.ErrNoRows {
			log.Println("error checking if row exists slots")
			return -1, err
    }
	if spotsLeft > 0 {
		return 0, nil
	}

	return 2, nil
}