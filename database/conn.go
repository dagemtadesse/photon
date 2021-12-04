package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	dbInstance    *sql.DB
	cacheInstance *redis.Client
	err           error

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

func GetCacheInstance() *redis.Client {
	if cacheInstance == nil {
		redis_db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

		cacheInstance = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			DB:       redis_db,
			Password: os.Getenv("REDIS_PASSWORD"),
		})
	}

	return cacheInstance
}
