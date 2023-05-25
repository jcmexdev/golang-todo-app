package mongo

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/models"
)

type UserRepository struct {
}

func (UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepository) Update(ctx context.Context, user models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepository) FindById(ctx context.Context, id interface{}) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepository) Close() error {
	//TODO implement me
	panic("implement me")
}
