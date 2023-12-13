package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wtran29/golang-restclient/rest"
)

var baseURL = "https://api.bookstore.com"

func TestMain(m *testing.M) {
	fmt.Println("starting auth test cases...")
	rest.StartMockupServer()

	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		URL:          baseURL + "/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	// client.Timeout = 5 * time.Second
	user, loginErr := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, loginErr)
	assert.EqualValues(t, http.StatusInternalServerError, loginErr.Status)
	assert.EqualValues(t, "invalid restClient response for user login", loginErr.Message)
}

func TestLoginInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          baseURL + "/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": true}`,
	})

	repository := usersRepository{}

	// client.Timeout = 5 * time.Second
	user, loginErr := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, loginErr)
	assert.EqualValues(t, http.StatusInternalServerError, loginErr.Status)
	assert.EqualValues(t, "invalid error interface when logging in user", loginErr.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          baseURL + "/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid rest response for user login", "status": 404, "error": true}`,
	})

	repository := usersRepository{}

	// client.Timeout = 5 * time.Second
	user, loginErr := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, loginErr)
	assert.EqualValues(t, http.StatusNotFound, loginErr.Status)
	assert.EqualValues(t, "invalid rest response for user login", loginErr.Message)
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          baseURL + "/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "54","first_name": "Test","last_name": "Dummy","email": "testdummy@example.com"}`,
	})

	repository := usersRepository{}

	user, loginErr := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, loginErr)
	assert.EqualValues(t, http.StatusInternalServerError, loginErr.Status)
	assert.EqualValues(t, "error when trying to unmarshal login response", loginErr.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          baseURL + "/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1,"first_name": "Test","last_name": "Dummy","email": "testdummy@example.com"}`,
	})

	repository := usersRepository{}

	user, loginErr := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, loginErr)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Test", user.FirstName)
	assert.EqualValues(t, "Dummy", user.LastName)
	assert.EqualValues(t, "testdummy@example.com", user.Email)

}
