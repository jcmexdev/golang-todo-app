package repository

import (
	"github.com/jxmexdev/go-todo-app/app/commons"
	"github.com/jxmexdev/go-todo-app/app/env"
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
