package task

import (
	"context"
	"fmt"
	taskDomain "tasoryx/internal/task/domain"
	"tasoryx/internal/task/port"
)

type service struct {
	repo port.Repo
}

func NewTask(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTask(ctx context.Context, task taskDomain.Task) (taskDomain.TaskID, error) {
	id, err := s.repo.Create(ctx, task)
	if err != nil {
		return id, fmt.Errorf("cannot create a task : %w", err)
	}
	return id, nil
}

func (s *service) UpdateTask(ctx context.Context, task taskDomain.Task, ID taskDomain.TaskID) error {
	if err := s.repo.Update(ctx, task, ID); err != nil {
		return fmt.Errorf("cannot update a task : %w", err)
	}
	return nil
}

func (s *service) GetTasks(ctx context.Context, filters ...taskDomain.FilterTask) ([]taskDomain.Task, error) {
	tasks, err := s.repo.Get(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("cannot get data : %w", err)
	}
	return tasks, nil
}

func (s *service) GetTaskByID(ctx context.Context, ID taskDomain.TaskID) (*taskDomain.Task, error) {
	task, err := s.repo.GetByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("cannot get task by ID : %w", err)
	}
	return task, nil
}

func (s *service) DeleteTask(ctx context.Context, ID taskDomain.TaskID) error {
	if err := s.repo.Delete(ctx, ID); err != nil {
		return fmt.Errorf("cannot delete task : %w", err)
	}
	return nil
}
