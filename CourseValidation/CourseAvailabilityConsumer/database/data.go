package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	secret "registerio/cv/consumer/secrets"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
}

type Consumer struct {
	Index              string
	RegisteredStudents []string
	MaxSize            int
	CurrentSize        int
}

func (s *DB) RetrieveState(index string) (Consumer, error) {
	retval := Consumer{}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}

	sql := `SELECT index, "max size" FROM "course availability" WHERE index=$1;`

	rows, err := db.Query(sql, index)
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&retval.Index, &retval.MaxSize)
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	sql = `SELECT ARRAY_AGG(netid), "class index"
	FROM "course registration" WHERE "class index" = $1 GROUP BY 2;`

	rows, err = db.Query(sql, index)
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(pq.Array(&retval.RegisteredStudents), &retval.Index)
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	retval.CurrentSize = len(retval.RegisteredStudents)

	return retval, nil

}

func (s *DB) AddRegistration(netID string, index string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}

	sqlStatement := `INSERT INTO "course registration" VALUES($1,$2);`

	_, err = db.Exec(sqlStatement, netID, index)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	return nil
}

func (s *DB) RemoveRegistration(netID string, index string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}

	sqlStatement := `DELETE FROM "course registration" WHERE "netid" = $1 AND "class index" = $2;`

	_, err = db.Exec(sqlStatement, netID, index)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	return nil
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
