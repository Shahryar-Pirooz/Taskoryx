package types

import (
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      uint8
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
