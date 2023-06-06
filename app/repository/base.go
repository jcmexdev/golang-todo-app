package repository

import "context"

type BaseRepository[T any] interface {
	Create(ctx context.Context, model *T) (*T, error)
	Update(ctx context.Context, model *T) (*T, error)
	FindById(ctx context.Context, id interface{}) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	Close() error
}

type CreateAction[T any] interface {
	Create(ctx context.Context, model *T) (*T, error)
}

type UpdateAction[T any] interface {
	Update(ctx context.Context, model *T) (*T, error)
}

type FindByIdAction[T any] interface {
	FindById(ctx context.Context, id interface{}) (*T, error)
}

type FindAllAction[T any] interface {
	FindAll(ctx context.Context) ([]*T, error)
}

type CloseAction interface {
	Close() error
}
