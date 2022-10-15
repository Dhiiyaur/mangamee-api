package bookmarkservice

import (
	"context"
)

type BookmarkService interface {
	GetBookmark(ctx context.Context, id string) (interface{}, error)
	InsertBookmark(ctx context.Context, bookmark interface{}) (string, error)
}
