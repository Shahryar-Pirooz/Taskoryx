package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserID = uuid.UUID

type UserRole int8

const (
	UserRoleUnknown UserRole = iota
	UserRoleAdmin
	UserRoleUser
)

type User struct {
	ID        UserID
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FilterUser struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

func (u UserRole) isValid() bool {
	return u == UserRoleAdmin || u == UserRoleUser
}

func passwordValidation(pass string) []string {
	var errs []string
	if len(pass) < 8 {
		errs = append(errs, "password must be at least 8 characters long")
	}
	if hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass); !hasUpper {
		errs = append(errs, "password must contain at least 1 uppercase letter (A-Z)")
	}
	if hasLower := regexp.MustCompile(`[a-z]`).MatchString(pass); !hasLower {
		errs = append(errs, "password must contain at least 1 lowercase letter (a-z)")
	}
	if hasNumberOrSpecial := regexp.MustCompile(`[0-9!@#$%^&*]`).MatchString(pass); !hasNumberOrSpecial {
		errs = append(errs, "password must contain at least 1 number (0-9) or special character (!@#$%^&*)")
	}
	return errs
}

func (u User) Validate() error {
	var errs []string

	if len(u.Name) < 1 {
		errs = append(errs, "name is required")
	}
	if len(u.Email) > 1 {
		if match, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, u.Email); !match {
			errs = append(errs, "invalid email format")
		}
	} else {
		errs = append(errs, "email is required")
	}

	if !u.Role.isValid() {
		errs = append(errs, fmt.Sprintf("status '%d' is invalid", u.Role))
	}

	pv := passwordValidation(u.Password)

	if len(pv) > 0 {
		errs = append(errs, pv...)
	}

	if len(errs) > 0 {
		return errors.New("validation failed: " + strings.Join(errs, ";"))
	}
	return nil
}
