package port

import (
	"context"
	"tasoryx/internal/task/domain"
)

type Repo interface {
	Create(ctx context.Context, task domain.Task) (domain.TaskID, error)
	GetByID(ctx context.Context, taskID domain.TaskID) (*domain.Task, error)
	Get(ctx context.Context, filter ...domain.FilterTask) ([]domain.Task, error)
	Update(ctx context.Context, task domain.Task, ID domain.TaskID) error
	Delete(ctx context.Context, ID domain.TaskID)
}
