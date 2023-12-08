package token

import (
	"strings"

	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
)

type Repository interface {
	GetTokenByID(string) (*Token, *errors.JsonError)
	CreateToken(Token) *errors.JsonError
	UpdateTokenExpiry(Token) *errors.JsonError
}

type Service interface {
	GetTokenByID(string) (*Token, *errors.JsonError)
	CreateToken(Token) *errors.JsonError
	UpdateTokenExpiry(Token) *errors.JsonError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetTokenByID implements Service.
func (s *service) GetTokenByID(tokenID string) (*Token, *errors.JsonError) {
	tokenID = strings.TrimSpace(tokenID)
	if len(tokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	token, err := s.repository.GetTokenByID(tokenID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *service) CreateToken(t Token) *errors.JsonError {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.repository.CreateToken(t)
}

func (s *service) UpdateTokenExpiry(t Token) *errors.JsonError {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateTokenExpiry(t)
}
