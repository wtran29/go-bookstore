package resterr

import (
	"errors"
	"net/http"
)

type JsonError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   bool          `json:"error"`
	Reason  []interface{} `json:"reason"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(msg string, err error) *JsonError {
	result := &JsonError{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   true,
	}
	if err != nil {
		result.Reason = append(result.Reason, err.Error())
	}
	return result
}

func NewNotFoundError(msg string, err error) *JsonError {
	result := &JsonError{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   true,
	}
	if err != nil {
		result.Reason = append(result.Reason, err.Error())
	}
	return result
}

func NewInternalServerError(msg string, err error) *JsonError {
	result := &JsonError{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   true,
	}
	if err != nil {
		result.Reason = append(result.Reason, err.Error())
	}
	return result
}

func NewUnauthorizedError(msg string, err error) *JsonError {
	result := &JsonError{
		Message: msg,
		Status:  http.StatusUnauthorized,
		Error:   true,
	}
	if err != nil {
		result.Reason = append(result.Reason, err.Error())
	}
	return result
}
