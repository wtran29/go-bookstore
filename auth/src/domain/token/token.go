package token

import (
	"fmt"
	"time"
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
