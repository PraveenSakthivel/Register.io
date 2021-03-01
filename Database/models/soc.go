package models 

import (
	"github.com/lib/pq"
	"database/sql"
	"fmt"
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

func RetrieveAllClasses(db *sql.DB) ([]Soc, error) {

	var query = "SELECT ca.\"max size\"-ca.\"current size\" as \"Available Slots\",s.location, s.level, s.school, s.department, s.class_number, s.index, s.name, s.section, s.meeting_location, s.meeting_times, s.exam, s.instructors, s.codes, s.synopsis, s.books"+
	" FROM public.socs s LEFT OUTER JOIN public.\"course availability\" ca ON ca.index = s.index;"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var classes []Soc
	for rows.Next() {
		var class Soc
		if err := rows.Scan(&class.Spots, &class.Location, &class.Level, &class.School, &class.Department, &class.ClassNumber, 
			&class.Index, &class.Name, &class.Section, &class.MeetingLocation, &class.MeetingTimes,
			&class.Exam, &class.Instructors, &class.Codes, &class.Synopsis, &class.Books); err != nil {
			return nil, err
		}
		classes = append(classes, class)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return classes, err
}