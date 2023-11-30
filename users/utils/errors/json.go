package errors

import "net/http"

type JsonError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
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
