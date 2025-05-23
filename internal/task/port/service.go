package port

import (
	"context"
	"tasoryx/internal/task/domain"
)

type Service interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.TaskID, error)
	UpdateTask(ctx context.Context, task domain.Task, ID domain.TaskID) error
	GetTasks(ctx context.Context, filters ...domain.FilterTask) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, ID domain.TaskID) (*domain.Task, error)
	DeleteTask(ctx context.Context, ID domain.TaskID) error
}
