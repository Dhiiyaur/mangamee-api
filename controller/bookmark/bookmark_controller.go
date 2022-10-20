package bookmarkcontroller

import "context"

type BookmarkController interface {
	InsertBookmark(c context.Context) error
	GetBookmark(c context.Context) error
}
