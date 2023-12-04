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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.JsonError) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.UpdateUser(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.JsonError {
	currentUser := &users.User{ID: userId}
	return currentUser.DeleteUser()
}
