package link_repository

import (
	"context"
	"mangamee-api/entity"
)

type LinkRepository interface {
	Insert(ctx context.Context, data entity.LinkRepository) error
	FindById(ctx context.Context, data entity.LinkRepository) (entity.LinkRepository, error)
}
