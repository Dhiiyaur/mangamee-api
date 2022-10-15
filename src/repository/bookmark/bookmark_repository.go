package bookmark_repository

import (
	"context"
	"mangamee-api/src/entity"
)

type BookmarkRepository interface {
	Insert(ctx context.Context, data entity.BookmarkRepository) error
	FindById(ctx context.Context, data entity.BookmarkRepository) (entity.BookmarkRepository, error)
}
