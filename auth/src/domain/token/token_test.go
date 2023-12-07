package token

import (
	"testing"
	"time"
)

func TestTokenConstants(t *testing.T) {
	if expiryTime != 24 {
		t.Error("expiry time should be 24 hours")
	}
}

func TestGetAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}

	if at.UserID != 0 {
		t.Error("new access token should not have associated user id")
	}
}

func TestIsExpired(t *testing.T) {
	at := Token{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expiry = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token expiry of three hours from now should not be expired")
	}
}
