// services package will handle all the business logic
package services

import (
	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.JsonError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.JsonError) {
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	result := &users.User{ID: userId}
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	return result, nil
}
