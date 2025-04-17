package mapper

import (
	"tasoryx/internal/user/domain"
	"tasoryx/pkg/adapters/storage/types"
)

func UserDomain2Repo(d domain.User) *types.User {
	return &types.User{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		Password:  d.Password,
		Role:      uint8(d.Role),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func UserRepo2Domain(r types.User) *domain.User {
	return &domain.User{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Password:  r.Password,
		Role:      domain.UserRole(r.Role),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
