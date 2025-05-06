package port

import (
	"context"
	"tasoryx/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, user domain.User, ID domain.UserID) error
	GetUserInfo(ctx context.Context, ID domain.UserID) (*domain.User, error)
	GetUsers(ctx context.Context, filters ...domain.FilterUser) ([]domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}
