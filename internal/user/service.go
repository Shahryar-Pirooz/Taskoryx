package user

import (
	"context"
	"fmt"
	"tasoryx/internal/user/domain"
	"tasoryx/internal/user/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return "", fmt.Errorf("cannot create user : %w", err)
	}
	return id, nil
}
func (s *service) UpdateUser(ctx context.Context, user domain.User, ID domain.UserID) error {
	if err := s.repo.Update(ctx, user, ID); err != nil {
		return fmt.Errorf("cannot update user : %w", err)
	}
	return nil
}
func (s *service) GetUserInfo(ctx context.Context, ID domain.UserID) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("cannot get data : %w", err)
	}
	return user, nil
}
func (s *service) GetUsers(ctx context.Context, filters ...domain.FilterUser) ([]domain.User, error) {
	users, err := s.repo.Get(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("cannot get data : %w", err)
	}
	return users, nil
}
