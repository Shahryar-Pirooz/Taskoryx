package storage

import (
	"context"
	"errors"
	"tasoryx/internal/task/domain"
	"tasoryx/internal/task/port"
	"tasoryx/pkg/adapters/storage/mapper"
	"tasoryx/pkg/adapters/storage/types"
	"tasoryx/pkg/fp"
	"time"

	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) port.Repo {
	return &taskRepo{db: db}
}

func (tr *taskRepo) Create(ctx context.Context, data domain.Task) (domain.TaskID, error) {
	task := mapper.TaskDomain2Repo(data)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	result := tr.db.WithContext(ctx).Create(task)
	if result.Error != nil {
		return task.ID, errors.New("failed to create task : " + result.Error.Error())
	}
	return task.ID, nil
}
func (tr *taskRepo) GetByID(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	task := new(types.Task)
	result := tr.db.WithContext(ctx).First(task, "id = ?", taskID)
	taskDomain := mapper.TaskRepo2Domain(*task)
	if result.Error != nil {
		return taskDomain, errors.New("failed to get task by id : " + result.Error.Error())
	}
	return taskDomain, nil
}
func (tr *taskRepo) Get(ctx context.Context, filter ...domain.FilterTask) ([]domain.Task, error) {
	var tasks []types.Task
	result := tr.db.WithContext(ctx).Find(&tasks, filter)
	if result.Error != nil {
		return nil, errors.New("failed to get tasks : " + result.Error.Error())
	}
	tasksDomain := fp.Map(tasks, mapper.TaskRepo2Domain)
	return tasksDomain, nil
}
func (tr *taskRepo) Update(ctx context.Context, NewRecord domain.Task, ID domain.TaskID) error {
	task, err := tr.GetByID(ctx, ID)
	if err != nil {
		return errors.New("failed to get task by id: " + err.Error())
	}
	newTask := mapper.TaskDomain2Repo(NewRecord)
	newTask.UpdatedAt = time.Now()
	result := tr.db.WithContext(ctx).Model(task).Updates(newTask)
	if result.Error != nil {
		return errors.New("failed to update task: " + result.Error.Error())
	}
	return nil
}
func (tr *taskRepo) Delete(ctx context.Context, ID domain.TaskID) error {
	result := tr.db.WithContext(ctx).Delete(&types.Task{}, ID)
	if result.Error != nil {
		return errors.New("failed to update task: " + result.Error.Error())
	}
	return nil
}
