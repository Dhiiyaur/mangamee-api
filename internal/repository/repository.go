package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"mangamee-api/internal/config"
	"mangamee-api/internal/entity"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	ctx = context.Background()
)

type Repository struct {
	db    *sql.DB
	redis *redis.Client
}

func New(db *sql.DB, redis *redis.Client) *Repository {
	return &Repository{
		db:    db,
		redis: redis,
	}
}

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

func createRedisKey(params entity.MangaParams) string {

	var redisKey string
	switch params.Path {
	case "index":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.PageNumber)
	case "detail":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.MangaId)
	case "search":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.Search)
	case "read":
		redisKey = fmt.Sprintf("%s-%s-%s-%s", params.Path, params.Source, params.MangaId, params.ChapterId)
	case "chapter":
		redisKey = fmt.Sprintf("%s-%s-%s", params.Path, params.Source, params.MangaId)
	}
	return redisKey
}

func (repo *Repository) InsertStatistic(params entity.MangaParams) error {

	_, err := repo.db.Exec(`INSERT INTO logs (api, source, title, chapter) VALUES ($1, $2, $3, $4)`, params.Path, params.Source, params.MangaId, params.ChapterId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) InsertCache(params entity.MangaParams, mangaData interface{}) error {

	expired := time.Duration(config.Cfg.Redis.Expired) * time.Second
	key := createRedisKey(params)

	val, _ := json.Marshal(mangaData)
	err := repo.redis.Set(ctx, key, val, expired).Err()

	if err != nil {
		return errors.New("error set cache")
	}
	return nil
}

func (repo *Repository) GetCache(params entity.MangaParams) (interface{}, error) {

	var returnArrData []entity.MangaData
	var returnData entity.MangaData

	key := createRedisKey(params)
	val, err := repo.redis.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, errors.New("key does not exist")
	case err != nil:
		return nil, errors.New("get failed")
	case val == "":
		return nil, errors.New("value is empty")
	}

	if params.Path == "index" || params.Path == "search" {
		err = json.Unmarshal([]byte(val), &returnArrData)
		if err != nil {
			return returnArrData, errors.New("cannot unmarshall data")
		}
		return returnArrData, nil

	} else if params.Path == "detail" || params.Path == "read" || params.Path == "chapter" {
		err = json.Unmarshal([]byte(val), &returnData)
		if err != nil {
			return returnData, errors.New("cannot unmarshall data")
		}
		return returnData, nil
	}

	return nil, errors.New("bad request")
}

func (repo *Repository) InsertLink(key string, longUrl string) error {

	expired := time.Duration(config.Cfg.Redis.Expired) * time.Second
	err := repo.redis.Set(ctx, key, longUrl, expired).Err()

	if err != nil {
		return errors.New("error set cache")
	}
	return nil
}

func (repo *Repository) GetLink(key string) (interface{}, error) {

	val, err := repo.redis.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, errors.New("key does not exist")
	case err != nil:
		return nil, errors.New("get failed")
	case val == "":
		return nil, errors.New("value is empty")
	}
	return val, nil
}

func (repo *Repository) InsertBookmark(key string, bookmark interface{}) error {

	expired := time.Duration(config.Cfg.Redis.Expired) * time.Second
	val, _ := json.Marshal(bookmark)
	err := repo.redis.Set(ctx, key, val, expired).Err()

	if err != nil {
		return errors.New("error set cache")
	}
	return nil
}

func (repo *Repository) GetBookmark(key string) (interface{}, error) {

	var bookmarkData interface{}

	val, err := repo.redis.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, errors.New("key does not exist")
	case err != nil:
		return nil, errors.New("get failed")
	case val == "":
		return nil, errors.New("value is empty")
	}

	err = json.Unmarshal([]byte(val), &bookmarkData)
	if err != nil {
		return nil, errors.New("cannot unmarshall data")
	}
	return bookmarkData, nil
}
