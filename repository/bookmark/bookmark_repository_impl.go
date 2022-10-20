package bookmark_repository

import (
	"context"
	"encoding/json"
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"time"

	"github.com/go-redis/redis/v8"
)

type BookmarkRepositoryImpl struct {
	redis *redis.Client
}

func NewBookmarkRepository(repository *redis.Client) BookmarkRepository {
	return &BookmarkRepositoryImpl{
		redis: repository,
	}
}

func (repository *BookmarkRepositoryImpl) Insert(ctx context.Context, bookmark entity.BookmarkRepository) error {

	expired := time.Duration(helper.GetExpiredTimeRedis()) * time.Minute

	val, _ := json.Marshal(bookmark.Data)
	err := repository.redis.Set(ctx, bookmark.Key, val, expired).Err()

	if err != nil {
		return exception.NewInternal()
	}
	return nil
}

func (repository *BookmarkRepositoryImpl) FindById(ctx context.Context, bookmark entity.BookmarkRepository) (entity.BookmarkRepository, error) {

	var returnData entity.BookmarkRepository
	result, err := repository.redis.Get(ctx, bookmark.Key).Result()

	switch {
	case err == redis.Nil:
		return returnData, exception.NewNotFound("manga.Key", returnData.Key)
	case err != nil:
		return returnData, exception.NewInternal()
	case result == "":
		return returnData, exception.NewInternal()
	}

	err = json.Unmarshal([]byte(result), &returnData.Data)
	if err != nil {
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}
