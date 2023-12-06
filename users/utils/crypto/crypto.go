package crypto

import (
	"crypto/md5"
	"encoding/hex"
	stdError "errors"

	"github.com/wtran29/go-bookstore/users/domain/users"
	"golang.org/x/crypto/bcrypt"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func PasswordMatches(plainText string, user users.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case stdError.Is(err, bcrypt.ErrMismatchedHashAndPassword):

			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
