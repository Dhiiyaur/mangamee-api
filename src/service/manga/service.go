package mangaservice

import (
	"context"
	"mangamee-api/src/entity"
)

type MangaService interface {
	GetIndex(ctx context.Context, params entity.MangaParams) ([]entity.MangaData, error)
	GetSearch(ctx context.Context, params entity.MangaParams) ([]entity.MangaData, error)
	GetDetail(ctx context.Context, params entity.MangaParams) (entity.MangaData, error)
	GetImage(ctx context.Context, params entity.MangaParams) (entity.MangaData, error)
	GetChapter(ctx context.Context, params entity.MangaParams) (entity.MangaData, error)
	GetMeta(ctx context.Context, params entity.MangaParams) (entity.MangaData, error)
}
