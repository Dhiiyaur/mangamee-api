package db

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func CreateConnection() *sql.DB {

	db, err := sql.Open("postgres", viper.GetString("POSTGRES_URL"))
	if err != nil {
		log.Println(err)
	}

	return db
}

func CreateRedisConnection() *redis.Client {

	opt, err := redis.ParseURL(viper.GetString("REDIS_URI"))
	if err != nil {
		log.Println(err)
	}

	return redis.NewClient(opt)
}
