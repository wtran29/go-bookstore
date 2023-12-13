package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	expiryTime                 = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type Token struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id,omitempty"`
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

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (t *TokenRequest) Validate() *errors.JsonError {
	switch t.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant type parameter")
	}

	// TODO: Validate parameters for each grant_type

	return nil
}

func GetNewAccessToken(userID int64) Token {
	return Token{
		UserID: userID,
		Expiry: time.Now().UTC().Add(expiryTime * time.Hour).Unix(),
	}
}

func (t Token) IsExpired() bool {
	now := time.Now().UTC()
	expiryTime := time.Unix(t.Expiry, 0)
	fmt.Println(expiryTime)
	return expiryTime.Before(now)
}

func (t *Token) GenerateToken() *errors.JsonError {
	token := fmt.Sprintf("at-%d-%d-ran", t.UserID, t.Expiry)
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error generating bcrypt hash for access token: %v", err.Error()))
	}
	t.AccessToken = string(hashedToken)

	return nil
}
