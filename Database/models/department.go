package models 

import (
	"database/sql"
	"log"
	"fmt"
	data "registerio/db/database"
)

func RetrieveDepartments(s *data.DB) (map[int32]string, error) {
	
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

	var query = "SELECT d.department, d.name FROM public.department_lookups d;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	departments := make(map[int32]string)
	for rows.Next() {
		var department int
		var name string
		if err := rows.Scan(&department, &name); err != nil {
			return nil, err
		}
		deptnum := fmt.Sprintf("%03d - %s", department, name)
		departments[int32(department)] = deptnum
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, err
}