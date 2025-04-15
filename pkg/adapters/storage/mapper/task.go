package mapper

import (
	"tasoryx/internal/task/domain"
	"tasoryx/pkg/adapters/storage/types"

	"github.com/google/uuid"
)

func TaskDomain2Repo(d domain.Task) types.Task {
	id := d.ID.String()
	return types.Task{
		ID:          id,
		Title:       d.Title,
		Description: d.Description,
		Status:      uint8(d.Status),
		DueDate:     d.DueDate,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func TaskRepo2Domain(r types.Task) domain.Task {
	id, _ := uuid.Parse(r.ID)
	return domain.Task{
		ID:          id,
		Title:       r.Title,
		Description: r.Description,
		Status:      domain.StatusTask(r.Status),
		DueDate:     r.DueDate,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
