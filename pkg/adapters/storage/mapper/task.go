package mapper

import (
	"tasoryx/internal/task/domain"
	"tasoryx/pkg/adapters/storage/types"
)

func TaskDomain2Repo(d domain.Task) *types.Task {
	return &types.Task{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Status:      uint8(d.Status),
		DueDate:     d.DueDate,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func TaskRepo2Domain(r types.Task) *domain.Task {
	return &domain.Task{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Status:      domain.StatusTask(r.Status),
		DueDate:     r.DueDate,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
