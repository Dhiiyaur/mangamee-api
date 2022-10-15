package shortener_repository

import (
	"context"
	"mangamee-api/src/entity"
)

type ShortenerRepository interface {
	Insert(ctx context.Context, data entity.ShortenerRepository) error
	FindById(ctx context.Context, data entity.ShortenerRepository) (entity.ShortenerRepository, error)
}
