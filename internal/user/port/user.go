package port

import (
	"context"
	"tasoryx/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByID(ctx context.Context, UserID domain.UserID) (*domain.User, error)
	Get(ctx context.Context, filter ...domain.FilterUser) ([]domain.User, error)
	Update(ctx context.Context, user domain.User, ID domain.UserID) error
	Delete(ctx context.Context, ID domain.UserID) error
}
