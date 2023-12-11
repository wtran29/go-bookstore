package errors

import (
	"errors"
	"net/http"
)

type JsonError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(msg string) *JsonError {
	return &JsonError{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   true,
	}
}

func NewNotFoundError(msg string) *JsonError {
	return &JsonError{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   true,
	}
}

func NewInternalServerError(msg string) *JsonError {
	return &JsonError{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   true,
	}
}

func NewUnauthorizedError(msg string) *JsonError {
	return &JsonError{
		Message: msg,
		Status:  http.StatusUnauthorized,
		Error:   true,
	}
}
