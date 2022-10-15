package shortenercontroller

import "context"

type ShortenerController interface {
	InsertLink(c context.Context) error
	GetLink(c context.Context) error
}
