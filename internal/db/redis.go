package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mangamee-api/internal/models"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	ctx = context.Background()
)

func createKey(route string, queryParams models.QueryParams) (string, string) {

	var redisKey string
	switch route {
	case "index":
		redisKey = fmt.Sprintf("%s-%s-%s", route, queryParams.Source, queryParams.Page)
	case "detail":
		redisKey = fmt.Sprintf("%s-%s-%s", route, queryParams.Source, queryParams.Id)
	case "search":
		redisKey = fmt.Sprintf("%s-%s-%s", route, queryParams.Source, queryParams.Search)
	case "read":
		redisKey = fmt.Sprintf("%s-%s-%s-%s", route, queryParams.Source, queryParams.Id, queryParams.ChapterId)
	case "chapter":
		redisKey = fmt.Sprintf("%s-%s-%s", route, queryParams.Source, queryParams.Id)
	}
	return redisKey, route

}

func CacheChecking(route string, queryParams models.QueryParams) (models.ReturnData, error) {

	MangaData := models.ReturnData{}
	redisKey, route := createKey(route, queryParams)

	rdb := CreateRedisConnection()
	val, err := rdb.Get(ctx, redisKey).Result()

	switch {
	case err == redis.Nil:
		return MangaData, errors.New("key does not exist")
	case err != nil:
		return MangaData, errors.New("get failed")
	case val == "":
		return MangaData, errors.New("value is empty")
	}

	if route == "index" || route == "search" {
		err = json.Unmarshal([]byte(val), &MangaData.Datas)
		if err != nil {
			return MangaData, errors.New("cannot unmarshall data")
		}
	} else if route == "detail" || route == "read" || route == "chapter" {

		err = json.Unmarshal([]byte(val), &MangaData.Data)
		if err != nil {
			return MangaData, errors.New("cannot unmarshall data")
		}
	}

	return MangaData, nil
}

func SetCache(route string, queryParams models.QueryParams, mangaData models.ReturnData) {

	ttl := time.Duration(viper.GetInt("CACHE_TIME")) * time.Hour
	redisKey, route := createKey(route, queryParams)

	var value interface{}

	if route == "index" || route == "search" {
		value, _ = json.Marshal(mangaData.Datas)
	} else if route == "detail" || route == "read" || route == "chapter" {
		value, _ = json.Marshal(mangaData.Data)
	}

	rdb := CreateRedisConnection()
	err := rdb.Set(ctx, redisKey, value, ttl).Err()
	if err != nil {
		log.Println(err)
	}

}
