package repository

import (
	"github.com/jxmexdev/go-todo-app/app/models"
)

type UserRepository interface {
	BaseRepository[models.User]
}
