// services package will handle all the business logic
package services

import (
	"fmt"

	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

var (
	UsersService usersRepository = &usersService{}
)

type usersService struct{}

type usersRepository interface {
	CreateUser(users.User) (*users.User, *errors.JsonError)
	GetUser(int64) (*users.User, *errors.JsonError)
	UpdateUser(bool, users.User) (*users.User, *errors.JsonError)
	DeleteUser(int64) *errors.JsonError
	SearchUser(string) (users.Users, *errors.JsonError)
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.JsonError) {
	fmt.Println(user)
	// if valid, err := crypto.PasswordMatches(user.Password, user); err != nil || !valid {
	// 	return nil, errors.ParseError(err)
	// }
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// user.CreatedAt = date.GetNowDBFormat()

	user.Status = users.StatusActive
	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) GetUser(userId int64) (*users.User, *errors.JsonError) {
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	result := &users.User{ID: userId}
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.JsonError) {
	current, err := u.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	// current.UpdatedAt, _ = date.GetNowDBFormat()
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
		if user.Password != "" {
			current.Password = user.Password
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Password = user.Password
	}

	if err := current.UpdateUser(); err != nil {
		return nil, err
	}
	return current, nil
}

func (u *usersService) DeleteUser(userId int64) *errors.JsonError {
	currentUser := &users.User{ID: userId}
	return currentUser.DeleteUser()
}

func (u *usersService) SearchUser(status string) (users.Users, *errors.JsonError) {
	user := &users.User{}
	return user.FindUserByStatus(status)

}
