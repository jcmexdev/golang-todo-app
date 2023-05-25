package repository

import "context"

type BaseRepository[T any] interface {
	Create(ctx context.Context, model T) (T, error)
	Update(ctx context.Context, model T) (T, error)
	FindById(ctx context.Context, id interface{}) (T, error)
	FindAll(ctx context.Context) ([]T, error)
	Close() error
}
