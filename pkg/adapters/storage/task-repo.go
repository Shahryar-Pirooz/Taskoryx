package storage

import (
	"context"
	taskDomain "tasoryx/internal/task/domain"
	taskPort "tasoryx/internal/task/port"

	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	DB *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) taskPort.Repo {
	return &taskRepo{DB: db}
}

func (tr *taskRepo) Create(ctx context.Context, task taskDomain.Task) (taskDomain.TaskID, error) {
	panic("")
}
func (tr *taskRepo) GetByID(ctx context.Context, taskID taskDomain.TaskID) (*taskDomain.Task, error) {
	panic("")
}
func (tr *taskRepo) Get(ctx context.Context, filter ...taskDomain.FilterTask) ([]taskDomain.Task, error) {
	panic("")
}
func (tr *taskRepo) Update(ctx context.Context, task taskDomain.Task, ID taskDomain.TaskID) error {
	panic("")
}
func (tr *taskRepo) Delete(ctx context.Context, ID taskDomain.TaskID) {
	panic("")
}
