package bookmark_repository

import (
	"context"
	"encoding/json"
	"errors"
	"mangamee-api/src/config"
	"mangamee-api/src/entity"
	"mangamee-api/src/exception"
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

	expired := time.Duration(config.Cfg.Redis.Expired) * time.Minute
	val, _ := json.Marshal(bookmark.Data)
	err := repository.redis.Set(ctx, bookmark.Key, val, expired).Err()

	if err != nil {
		return errors.New("error set cache")
	}
	return nil
}

func (repository *BookmarkRepositoryImpl) FindById(ctx context.Context, bookmark entity.BookmarkRepository) (entity.BookmarkRepository, error) {

	var returnData entity.BookmarkRepository
	result, err := repository.redis.Get(ctx, bookmark.Key).Result()

	switch {
	case err == redis.Nil:
		return returnData, exception.NewErrorMsg(exception.CodeErrDataNotFound, err)
	case err != nil:
		return returnData, exception.NewErrorMsg(exception.CodeInternalServerError, err)
	case result == "":
		return returnData, exception.NewErrorMsg(exception.CodeErrDataNotFound, err)
	}

	err = json.Unmarshal([]byte(result), &returnData.Data)
	if err != nil {
		return returnData, errors.New("cannot unmarshall data")
	}

	return returnData, nil
}
