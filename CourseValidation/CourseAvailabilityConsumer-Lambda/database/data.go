package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	consumerState "registerio/cv/consumer-lambda/protobuf"
	secret "registerio/cv/consumer-lambda/secrets"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/proto"
)

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
	Cache    string
}

type Consumer struct {
	Index              string
	RegisteredStudents []string
	MaxSize            int
	CurrentSize        int
}

func (s *DB) RetrieveState(index string) (Consumer, error) {
	log.Println("Checking Cache")
	mc := memcache.New(s.Cache)
	stateByte, err := mc.Get(index)
	if err != nil {
		log.Println("Retreiving State from DB: ", err)
		return s.RetrieveStateDB(index)
	}
	state := &consumerState.State{}
	if err = proto.Unmarshal(stateByte.Value, state); err != nil {
		log.Println("Error decoding memcached state: ", err)
		return s.RetrieveStateDB(index)
	}
	log.Println("Using Cache")
	finalState := Consumer{Index: index, RegisteredStudents: state.RegisteredStudents, MaxSize: int(state.MaxSize), CurrentSize: int(state.CurrentSize)}
	return finalState, nil

}

func (s *DB) PushState(consumer Consumer) {
	mc := memcache.New(s.Cache)
	state := &consumerState.State{
		MaxSize:            int32(consumer.MaxSize),
		CurrentSize:        int32(consumer.CurrentSize),
		RegisteredStudents: consumer.RegisteredStudents,
	}
	stateByte, err := proto.Marshal(state)
	if err != nil {
		log.Println("Error serializing state: ", err)
		return
	}
	err = mc.Set(&memcache.Item{Key: consumer.Index, Value: stateByte})
	if err != nil {
		log.Println("Error commiting to cache: ", err)
	}
	return
}

func (s *DB) RetrieveStateDB(index string) (Consumer, error) {
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

	sql = `SELECT ARRAY_AGG(netid), "class_index"
	FROM "course_registrations" WHERE "class_index" = $1 GROUP BY 2;`

	rows, err = db.Query(sql, index)
	if err != nil {
		log.Println("Database error: ", err)
		return retval, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(pq.Array(&retval.RegisteredStudents), &retval.Index)
		log.Println(retval.RegisteredStudents)
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return retval, err
		}

		err = rows.Err()
		if err != nil {
			log.Println("Error Parsing records: ", err)
			return retval, err
		}
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

	sqlStatement := `INSERT INTO "course_registrations" VALUES($1,$2);`

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

	sqlStatement := `DELETE FROM "course_registrations" WHERE "netid" = $1 AND "class_index" = $2;`

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
	cachestring, err := secret.GetTokenSecret("prod/CacheUrl")
	if err != nil {
		return nil, err
	}
	var url struct {
		Url string `json:"url"`
	}
	err = json.Unmarshal([]byte(cachestring), &url)
	if err != nil {
		return nil, err
	}
	retval.Cache = url.Url
	return &retval, nil
}
