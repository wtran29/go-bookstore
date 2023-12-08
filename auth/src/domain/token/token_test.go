package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenConstants(t *testing.T) {
	if expiryTime != 24 {
		t.Error("expiry time should be 24 hours")
	}
	assert.EqualValues(t, 24, expiryTime)
}

func TestGetAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "new access token should not be expired")
	if at.IsExpired() {
		t.Error("new access token should not be expired")
	}

	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	// if at.AccessToken != "" {
	// 	t.Error("new access token should not have defined access token id")
	// }

	assert.True(t, at.UserID == 0)
	// if at.UserID != 0 {
	// 	t.Error("new access token should not have associated user id")
	// }
}

func TestIsExpired(t *testing.T) {
	at := Token{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }
	at.Expiry = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiry of three hours from now should not be expired")
	// if at.IsExpired() {
	// 	t.Error("access token expiry of three hours from now should not be expired")
	// }
}
