package types

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
