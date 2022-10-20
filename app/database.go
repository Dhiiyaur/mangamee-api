package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mangamee-api/logger"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type dataSources struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

func InitDS() (*dataSources, error) {

	log.Printf("Initializing data sources\n")

	redisUri := os.Getenv("REDIS_URI")
	log.Printf("Connecting to Redis\n")
	opt, err := redis.ParseURL(redisUri)

	if err != nil {
		logger.Error("error opening redis:", zap.Error(err))
		return nil, fmt.Errorf("error opening redis: %w", err)
	}
	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		logger.Error("error connecting to redis:", zap.Error(err))
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}

	pgUri := os.Getenv("POSTGRES_URI")
	log.Printf("Connecting to Postgresql\n")
	db, err := sql.Open("postgres", pgUri)
	if err != nil {
		logger.Error("error opening db", zap.Error(err))
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		logger.Error("error connecting to db", zap.Error(err))
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return &dataSources{
		DB:          db,
		RedisClient: rdb,
	}, nil
}
