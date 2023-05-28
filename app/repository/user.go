package repository

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/models"
)

type UserRepository interface {
	BaseRepository[models.User]
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

func CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return implementations.UserRepository.Create(ctx, user)
}

func UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return implementations.UserRepository.Update(ctx, user)
}

func FindUserById(ctx context.Context, id interface{}) (*models.User, error) {
	return implementations.UserRepository.FindById(ctx, id)
}

func FindAllUsers(ctx context.Context) ([]*models.User, error) {
	return implementations.UserRepository.FindAll(ctx)
}

func CloseUser() error {
	return implementations.UserRepository.Close()
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementations.UserRepository.FindByEmail(ctx, email)
}
