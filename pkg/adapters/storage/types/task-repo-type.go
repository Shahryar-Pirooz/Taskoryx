package types

import (
	"time"
)

type Task struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string
	Description string
	Status      uint8
	DueDate     time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:nano"`
	DeletedAt   time.Time `gorm:"autoDeleteTime:nano"`
}
