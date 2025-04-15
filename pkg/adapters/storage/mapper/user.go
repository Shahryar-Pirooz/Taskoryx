package mapper

import (
	"tasoryx/internal/user/domain"
	"tasoryx/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

func UserDomain2Repo(d domain.User) types.User {
	id := d.ID.String()
	return types.User{
		ID:        id,
		Name:      d.Name,
		Email:     d.Email,
		Password:  d.Password,
		Role:      uint8(d.Role),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func UserRepo2Domain(r types.User) domain.User {
	id, _ := uuid.Parse(r.ID)
	return domain.User{
		ID:        id,
		Name:      r.Name,
		Email:     r.Email,
		Password:  r.Password,
		Role:      domain.UserRole(r.Role),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
