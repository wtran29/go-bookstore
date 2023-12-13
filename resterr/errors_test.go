package resterr

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, true, err.Error)

	assert.NotNil(t, err.Reason)
	assert.EqualValues(t, 1, len(err.Reason))
	assert.EqualValues(t, "database error", err.Reason[0])

}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message", errors.New("bad request"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, true, err.Error)

	assert.NotNil(t, err.Reason)
	assert.EqualValues(t, 1, len(err.Reason))
	assert.EqualValues(t, "bad request", err.Reason[0])
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message", errors.New("not found error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, true, err.Error)

	assert.NotNil(t, err.Reason)
	assert.EqualValues(t, 1, len(err.Reason))
	assert.EqualValues(t, "not found error", err.Reason[0])
}

func TestNewError(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"test message", "test message"},
		{"another message", "another message"},
	}

	// Iterate over test cases
	for _, tc := range tests {
		// Call the function with the test input
		err := NewError(tc.input)

		// Check if the returned error has the expected message
		if err.Error() != tc.want {
			t.Errorf("NewError(%s) = %s; want %s", tc.input, err.Error(), tc.want)
		}
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message", errors.New("not authorized error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, true, err.Error)

	assert.NotNil(t, err.Reason)
	assert.EqualValues(t, 1, len(err.Reason))
	assert.EqualValues(t, "not authorized error", err.Reason[0])
}
