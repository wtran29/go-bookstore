package users

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/wtran29/go-bookstore/resterr"
)

const (
	StatusActive = "active"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
	Password  string    `json:"password"`
}

type Users []User

func (user *User) Validate() *resterr.JsonError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resterr.NewBadRequestError("validation error", errors.New("email cannot be empty"))
	}
	if !isValidEmail(user.Email) {
		return resterr.NewBadRequestError("validation error", errors.New("invalid email address"))
	}

	if len(user.Password) == 0 {
		return resterr.NewBadRequestError("validation error", errors.New("password cannot be empty"))
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validEmail := regexp.MustCompile(emailRegex)
	return validEmail.MatchString(email)
}
