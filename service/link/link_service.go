package linkservice

import (
	"context"
	"mangamee-api/entity"
)

type LinkService interface {
	InsertLink(ctx context.Context, url string) (entity.LinkRepository, error)
	GetLink(ctx context.Context, id string) (entity.LinkRepository, error)
}
