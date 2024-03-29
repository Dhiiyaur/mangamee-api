package manga_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"

	"time"

	"github.com/go-redis/redis/v8"
)

type MangaRepositoryImpl struct {
	redis *redis.Client
	db    *sql.DB
}

func NewMangaRepository(redis *redis.Client, db *sql.DB) MangaRepository {
	return &MangaRepositoryImpl{
		redis: redis,
		db:    db,
	}
}

func (repository *MangaRepositoryImpl) InsertCache(ctx context.Context, manga entity.MangaRepository) error {

	// expired := time.Duration(config.Cfg.Redis.Expired) * time.Minute
	expired := time.Duration(helper.GetExpiredTimeRedis()) * time.Minute
	val, _ := json.Marshal(manga.Data)
	err := repository.redis.Set(ctx, manga.Key, val, expired).Err()

	if err != nil {
		return exception.NewInternal()
	}
	return nil
}

func (repository *MangaRepositoryImpl) FindByIdCache(ctx context.Context, manga entity.MangaRepository, path string) (entity.MangaRepository, error) {

	var arrayTypeData []entity.MangaData
	var singleTypeData entity.MangaData
	var mangaData entity.MangaRepository

	val, err := repository.redis.Get(ctx, manga.Key).Result()
	switch {
	case err == redis.Nil:
		return mangaData, exception.NewNotFound("manga.Key", manga.Key)
	case err != nil:
		return mangaData, exception.NewInternal()
	case val == "":
		return mangaData, exception.NewInternal()
	}

	if path == "index" || path == "search" {
		err = json.Unmarshal([]byte(val), &arrayTypeData)
		if err != nil {
			return mangaData, exception.NewInternal()

		}
		mangaData.Data = arrayTypeData
		return mangaData, nil

	} else if path == "detail" || path == "read" || path == "chapter" {
		err = json.Unmarshal([]byte(val), &singleTypeData)
		if err != nil {
			return mangaData, exception.NewInternal()
		}

		mangaData.Data = singleTypeData
		return mangaData, nil
	}

	return mangaData, nil
}

func (repository *MangaRepositoryImpl) InsertStatistic(ctx context.Context, manga entity.RequestParams) error {

	_, err := repository.db.Exec(`INSERT INTO logs (api, source, title, chapter) VALUES ($1, $2, $3, $4)`, manga.Path, manga.Source, manga.MangaId, manga.ChapterId)
	if err != nil {
		return exception.NewInternal()
	}
	return nil
}
