package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
)

const (
	expiryTime = 24
)

type Token struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expiry      int64  `json:"expiry"`
}

func (t *Token) Validate() *errors.JsonError {
	t.AccessToken = strings.TrimSpace(t.AccessToken)
	if t.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if t.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if t.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if t.Expiry <= 0 {
		return errors.NewBadRequestError("invalid token expiry")
	}
	return nil
}

func GetNewAccessToken() Token {
	return Token{
		Expiry: time.Now().UTC().Add(expiryTime * time.Hour).Unix(),
	}
}

func (t Token) IsExpired() bool {
	now := time.Now().UTC()
	expiryTime := time.Unix(t.Expiry, 0)
	fmt.Println(expiryTime)
	return expiryTime.Before(now)
}
