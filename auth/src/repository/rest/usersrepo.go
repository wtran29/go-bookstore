package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/wtran29/go-bookstore/auth/src/domain/users"

	"github.com/wtran29/go-bookstore/resterr"
	"github.com/wtran29/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "https://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *resterr.JsonError)
}

type usersRepository struct{}

func NewRepo() UsersRepository {
	return &usersRepository{}
}

func (u *usersRepository) LoginUser(email string, password string) (*users.User, *resterr.JsonError) {
	body := users.Login{
		Email:    email,
		Password: password,
	}

	resp := restClient.Post("/users/login", body)
	fmt.Println("Response:", resp)
	if resp == nil || resp.Body == nil {
		return nil, resterr.NewInternalServerError("response error", errors.New("invalid restClient response for user login"))
	}

	if resp.StatusCode > 299 {
		fmt.Println(resp)
		var restErr resterr.JsonError
		if err := json.Unmarshal(resp.Bytes(), &restErr); err != nil {
			return nil, resterr.NewInternalServerError("response error", errors.New("invalid error interface when logging in user"))
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Bytes(), &user); err != nil {
		return nil, resterr.NewInternalServerError("unmarshal error", errors.New("error when trying to unmarshal login response"))
	}
	return &user, nil
}
