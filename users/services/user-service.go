// services package will handle all the business logic
package services

import (
	"github.com/wtran29/go-bookstore/resterr"
	"github.com/wtran29/go-bookstore/users/domain/users"
	"golang.org/x/crypto/bcrypt"
)

var (
	UsersService usersRepository = &usersService{}
)

type usersService struct{}

type usersRepository interface {
	CreateUser(users.User) (*users.User, *resterr.JsonError)
	GetUser(int64) (*users.User, *resterr.JsonError)
	UpdateUser(bool, users.User) (*users.User, *resterr.JsonError)
	DeleteUser(int64) *resterr.JsonError
	SearchUser(string) (users.Users, *resterr.JsonError)
	LoginUser(users.Login) (*users.User, *resterr.JsonError)
}

func (u *usersService) CreateUser(user users.User) (*users.User, *resterr.JsonError) {

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

func (u *usersService) GetUser(userId int64) (*users.User, *resterr.JsonError) {
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	result := &users.User{ID: userId}
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *resterr.JsonError) {
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

func (u *usersService) DeleteUser(userId int64) *resterr.JsonError {
	currentUser := &users.User{ID: userId}
	return currentUser.DeleteUser()
}

func (u *usersService) SearchUser(status string) (users.Users, *resterr.JsonError) {
	user := &users.User{}
	return user.FindUserByStatus(status)

}

func (u *usersService) LoginUser(login users.Login) (*users.User, *resterr.JsonError) {

	user := &users.User{
		Email: login.Email,
	}
	user, err := user.FindByEmail()
	if err != nil {
		return nil, err
	}

	pwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if pwErr != nil {
		return nil, resterr.NewUnauthorizedError("Invalid credentials", pwErr)
	}

	return user, nil
}
