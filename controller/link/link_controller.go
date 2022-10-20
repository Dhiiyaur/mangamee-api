package linkcontroller

import "context"

type LinkController interface {
	InsertLink(c context.Context) error
	GetLink(c context.Context) error
}
