package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "database-1.cluster-cpecpwkhwaq9.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "registerio"
	password = "registera"
	dbname   = "maindb"
)

type Consumer struct {
	Index              string
	RegisteredStudents []string
	MaxSize            int
	CurrentSize        int
}

func RetrieveState(index string) (Consumer, error) {
	retval := Consumer{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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

	sql := `SELECT index, "registered students", "max size", "current size" FROM "course availability" WHERE index=$1;`

	rows, err := db.Query(sql, index)
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&retval.Index, pq.Array(&retval.RegisteredStudents), &retval.MaxSize, &retval.CurrentSize)
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error Parsing records: ", err)
		return retval, err
	}

	return retval, nil

}

func UpdateState(state Consumer) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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

	sqlStatement := `UPDATE "course availability" SET "registered students" = $1, "current size" = $2 WHERE index = $3;`

	_, err = db.Exec(sqlStatement, pq.Array(state.RegisteredStudents), state.CurrentSize, state.Index)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	return nil
}
