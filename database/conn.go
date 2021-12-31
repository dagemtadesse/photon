package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	err        error

	Ctx = context.Background()
)

func GetDBInstance() *sql.DB {
	if dbInstance == nil {
		db_url := os.Getenv("PG_URL")
		dbInstance, err = sql.Open("postgres", db_url)

		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("Connected to DB")
		}
	}

	return dbInstance
}

func RedisStore() redis.Store {
	store, err := redis.NewStore(
		10,
		os.Getenv("REDIS_NETWORK"),
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("HASHING_KEY")),
	)

	if err != nil {
		log.Fatalf("Unable to connect to redis: %v", err.Error())
	}

	return store
}
