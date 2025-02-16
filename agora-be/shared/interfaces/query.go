package interfaces

import "context"

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}
