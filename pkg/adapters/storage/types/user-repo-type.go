package types

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      uint8
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
	DeletedAt time.Time `gorm:"autoDeleteTime:nano"`
}
