package shortener_repository

import (
	"context"
	"mangamee-api/src/config"
	"mangamee-api/src/entity"
	"mangamee-api/src/exception"
	"time"

	"github.com/go-redis/redis/v8"
)

type ShortenerRepositoryImpl struct {
	redis *redis.Client
}

func NewShortenerRepository(repository *redis.Client) ShortenerRepository {
	return &ShortenerRepositoryImpl{
		redis: repository,
	}
}

func (repository *ShortenerRepositoryImpl) Insert(ctx context.Context, data entity.ShortenerRepository) error {

	expired := time.Duration(config.Cfg.Redis.Expired) * time.Minute
	err := repository.redis.Set(ctx, data.Key, data.LongUrl, expired).Err()

	if err != nil {
		return exception.NewErrorMsg(exception.CodeInternalServerError, err)
	}
	return nil
}

func (repository *ShortenerRepositoryImpl) FindById(ctx context.Context, data entity.ShortenerRepository) (entity.ShortenerRepository, error) {

	var returnData entity.ShortenerRepository

	result, err := repository.redis.Get(ctx, data.Key).Result()
	switch {
	case err == redis.Nil:
		return returnData, exception.NewErrorMsg(exception.CodeErrDataNotFound, err)
	case err != nil:
		return returnData, exception.NewErrorMsg(exception.CodeInternalServerError, err)
	case result == "":
		return returnData, exception.NewErrorMsg(exception.CodeErrDataNotFound, err)
	}

	returnData.LongUrl = result
	return returnData, nil
}
