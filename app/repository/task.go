package repository

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/models"
)

type TaskRepository interface {
	CreateAction[models.Task]
	UpdateAction[models.Task]
	FindByIdAction[models.Task]
	FindAllByUserId(ctx context.Context, userId string) ([]*models.Task, error)
	Delete(ctx context.Context, id string) error
}

func CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	return implementations.TaskRepository.Create(ctx, task)
}
func UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	return implementations.TaskRepository.Update(ctx, task)
}

func FindTaskById(ctx context.Context, id string) (*models.Task, error) {
	return implementations.TaskRepository.FindById(ctx, id)
}

func FindAllTasksByUserId(ctx context.Context, userId string) ([]*models.Task, error) {
	return implementations.TaskRepository.FindAllByUserId(ctx, userId)
}

func DeleteTask(ctx context.Context, id string) error {
	return implementations.TaskRepository.Delete(ctx, id)
}
