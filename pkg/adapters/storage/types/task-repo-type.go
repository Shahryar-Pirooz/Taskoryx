package types

import "time"

type Task struct {
	ID          uint8     `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      uint8     `db:"status"`
	DateDue     time.Time `db:"date_due"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
