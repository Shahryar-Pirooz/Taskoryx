package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type TaskID = string

type StatusTask int8

const (
	StatusTaskUnknown StatusTask = iota
	StatusTaskDone
	StatusTaskUnDone
)

func (s StatusTask) isValid() bool {
	return s == StatusTaskDone || s == StatusTaskUnDone
}

type Task struct {
	ID          TaskID
	Title       string
	Description string
	Status      StatusTask
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FilterTask struct {
	Title  string
	Status StatusTask
}

func (t Task) Validate() error {
	var errs []string

	if t.Title == "" {
		errs = append(errs, "title is required")
	}

	if !t.Status.isValid() {
		errs = append(errs, fmt.Sprintf("status '%d' is invalid", t.Status))
	}
	if t.DueDate.Before(time.Now()) {
		errs = append(errs, "due date must be in the future")
	}

	if len(errs) > 0 {
		return errors.New("validation failed: " + strings.Join(errs, ";"))
	}
	return nil
}
