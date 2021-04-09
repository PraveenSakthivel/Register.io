package models 

import (
	"github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
	data "registerio/db/database"
)

type Soc struct {
	Spots			int
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
	Instructors     pq.StringArray
	Codes           pq.StringArray
	Synopsis        string
	Books           pq.StringArray
}

func RetrieveAllClasses(s *data.DB) (map[string][]Soc, error) {
	
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

	var query = "SELECT cx.slots, s.location, s.level, s.school, s.department, s.\"class number\", "+
	"s.index, s.name, s.section, s.\"meeting location\", s.\"meeting times\", s.exam, s.instructors, s.codes, "+
	"s.synopsis, s.books "+ 
	"FROM public.soc s LEFT OUTER JOIN ("+
	"SELECT ca.index, (ca.\"max size\" - COALESCE(crr.amt_filled,0)) as \"slots\" FROM \"course availability\" ca LEFT OUTER JOIN (SELECT cr.class_index, COUNT(cr.netid) as \"amt_filled\" from course_registrations cr group by class_index) crr "+
	"ON ca.index = crr.class_index) cx "+
	"ON s.index = cx.index;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	classes := make(map[string][]Soc)
	for rows.Next() {
		var class Soc
		if err := rows.Scan(&class.Spots, &class.Location, &class.Level, &class.School, &class.Department, &class.ClassNumber, 
			&class.Index, &class.Name, &class.Section, &class.MeetingLocation, &class.MeetingTimes,
			&class.Exam, &class.Instructors, &class.Codes, &class.Synopsis, &class.Books); err != nil {
			return nil, err
		}
		coursenumber := fmt.Sprintf("%02d:%03d:%03d", class.School, class.Department, class.ClassNumber)
		classes[coursenumber] = append(classes[coursenumber], class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return classes, err
}