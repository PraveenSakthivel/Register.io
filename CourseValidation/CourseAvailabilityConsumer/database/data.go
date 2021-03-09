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

func AddRegistration(netID string, index string) error {
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

	sqlStatement := `INSERT INTO "course registration" VALUES($1,$2);`

	_, err = db.Exec(sqlStatement, netID, index)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	return nil
}

func RemoveRegistration(netID string, index string) error {
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

	sqlStatement := `DELETE FROM "course registration" WHERE "netid" = $1 AND "class index" = $2;`

	_, err = db.Exec(sqlStatement, netID, index)
	if err != nil {
		log.Println("Database error: ", err)
		return err
	}
	return nil
}
