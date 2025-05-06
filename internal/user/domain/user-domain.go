package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type UserID = string

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
	Name  string
	Email string
	Role  UserRole
}

func (u UserRole) isValid() bool {
	return u == UserRoleAdmin || u == UserRoleUser
}

func PasswordValidation(pass string) []string {
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

func EmailValidation(email string) []string {
	var errs []string
	if len(email) > 1 {
		if match, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email); !match {
			errs = append(errs, "invalid email format")
		}
	} else {
		errs = append(errs, "email is required")
	}
	return errs
}

func (u User) Validate() error {
	var errs []string

	if len(u.Name) < 1 {
		errs = append(errs, "name is required")
	}

	ev := EmailValidation(u.Email)

	if len(ev) > 0 {
		errs = append(errs, ev...)
	}

	if !u.Role.isValid() {
		errs = append(errs, fmt.Sprintf("status '%d' is invalid", u.Role))
	}

	pv := PasswordValidation(u.Password)

	if len(pv) > 0 {
		errs = append(errs, pv...)
	}

	if len(errs) > 0 {
		return errors.New("validation failed: " + strings.Join(errs, ";"))
	}
	return nil
}
