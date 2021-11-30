package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	err        error
)

func GetInstance() *sql.DB {
	if dbInstance == nil {
		db_url := os.Getenv("PG_URL")
		dbInstance, err = sql.Open("postgres", db_url)

		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("Connected to DB")
		}

		return dbInstance
	}

	return dbInstance
}
