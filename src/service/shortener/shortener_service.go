package linkshortenerservice

import (
	"context"
	"mangamee-api/src/entity"
)

type ShortenerService interface {
	InsertLink(ctx context.Context, url string) (entity.ShortenerRepository, error)
	GetLink(ctx context.Context, id string) (entity.ShortenerRepository, error)
}
