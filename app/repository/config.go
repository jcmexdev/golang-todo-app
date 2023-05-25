package repository

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/commons"
	"github.com/jxmexdev/go-todo-app/app/env"
	"github.com/jxmexdev/go-todo-app/app/models"
	"github.com/jxmexdev/go-todo-app/app/repository/mongo"
)

type Repositories struct {
	UserRepository UserRepository
}

var implementations Repositories

func LoadConfiguration() {
	switch env.Conf.DbDriver {
	case commons.MongoDbDriver:
		implementations = Repositories{
			UserRepository: mongo.NewUserMongoRepository(),
		}
	default:
		panic("Invalid db driver for repository implementation")
	}
}

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return implementations.UserRepository.Create(ctx, user)
}

func UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return implementations.UserRepository.Update(ctx, user)
}

func FindUserById(ctx context.Context, id interface{}) (models.User, error) {
	return implementations.UserRepository.FindById(ctx, id)
}

func FindAllUsers(ctx context.Context) ([]models.User, error) {
	return implementations.UserRepository.FindAll(ctx)
}

func CloseUser() error {
	return implementations.UserRepository.Close()
}
