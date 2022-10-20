package link_repository

import (
	"context"

	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"time"

	"github.com/go-redis/redis/v8"
)

type LinkRepositoryImpl struct {
	redis *redis.Client
}

func NewLinkRepository(repository *redis.Client) LinkRepository {
	return &LinkRepositoryImpl{
		redis: repository,
	}
}

func (repository *LinkRepositoryImpl) Insert(ctx context.Context, data entity.LinkRepository) error {

	expired := time.Duration(helper.GetExpiredTimeRedis()) * time.Minute
	err := repository.redis.Set(ctx, data.Key, data.LongUrl, expired).Err()

	if err != nil {
		return exception.NewInternal()
	}
	return nil
}

func (repository *LinkRepositoryImpl) FindById(ctx context.Context, data entity.LinkRepository) (entity.LinkRepository, error) {

	var returnData entity.LinkRepository

	result, err := repository.redis.Get(ctx, data.Key).Result()
	switch {
	case err == redis.Nil:
		return returnData, exception.NewNotFound("Link.Key", returnData.Key)
	case err != nil:
		return returnData, exception.NewInternal()
	case result == "":
		return returnData, exception.NewInternal()
	}

	returnData.LongUrl = result
	return returnData, nil
}
