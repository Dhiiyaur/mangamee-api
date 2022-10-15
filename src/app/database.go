package app

import (
	"database/sql"
	"mangamee-api/src/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func CreateDbConnection(cfg config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.Database.URI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateRedisConnection(cfg config.Config) (*redis.Client, error) {

	opt, err := redis.ParseURL(cfg.Redis.URI)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opt), nil
}
