package manga_repository

import (
	"context"
	"mangamee-api/src/entity"
)

type MangaRepository interface {
	InsertCache(ctx context.Context, manga entity.MangaRepository) error
	FindByIdCache(ctx context.Context, manga entity.MangaRepository, path string) (entity.MangaRepository, error)
	InsertStatistic(ctx context.Context, manga entity.MangaParams) error
}
