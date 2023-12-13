package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
	"github.com/wtran29/golang-restclient/rest"
)

const (
	headerXPublic    = "X-Public"
	headerXClientId  = "X-Client-Id"
	headerXCallerId  = "X-Caller-Id"
	paramAccessToken = "access_token"
)

var (
	authRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	ID       string `json:"id"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
}

type authInterface struct {
}

func IsPublic(r *http.Request) bool {
	if r == nil {
		return true
	}
	return r.Header.Get(headerXPublic) == "true"
}

func GetCallerId(r *http.Request) int64 {
	if r == nil {
		return 0
	}
	callerId, err := strconv.ParseInt(r.Header.Get(headerXCallerId), 10, 64)
	if err != nil {
		return 0
	}
	return callerId
}

func GetClientId(r *http.Request) int64 {
	if r == nil {
		return 0
	}
	clientId, err := strconv.ParseInt(r.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}
	return clientId
}

func AuthenticateRequest(r *http.Request) *errors.JsonError {
	if r == nil {
		return nil
	}

	cleanRequest(r)

	tokenID := strings.TrimSpace(r.URL.Query().Get(paramAccessToken))
	if tokenID == "" {
		return nil
	}
	at, err := getAccessToken(tokenID)
	fmt.Println(err)
	if err != nil {
		if err.Status == http.StatusNotFound {
			return nil
		}
		return err
	}

	r.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientID))
	r.Header.Add(headerXCallerId, fmt.Sprintf("%v", at.UserID))

	return nil
}

func cleanRequest(r *http.Request) {
	if r == nil {
		return
	}
	r.Header.Del(headerXClientId)
	r.Header.Del(headerXCallerId)

}

func getAccessToken(tokenID string) (*accessToken, *errors.JsonError) {
	resp := authRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", tokenID))
	if resp == nil || resp.Body == nil {
		return nil, errors.NewInternalServerError("invalid restClient response retrieving access token")
	}

	if resp.StatusCode > 299 {
		fmt.Println(resp)
		var restErr errors.JsonError
		if err := json.Unmarshal(resp.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid error interface retreiving access token")
		}

		return nil, &restErr
	}

	var at accessToken
	if err := json.Unmarshal(resp.Bytes(), &at); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return &at, nil
}
