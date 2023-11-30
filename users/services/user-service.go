// services package will handle all the business logic
package services

import (
	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.JsonError) {
	return &user, nil
}
