package services

import (
	"errors"
	"strings"

	"github.com/wtran29/go-bookstore/auth/src/domain/token"
	"github.com/wtran29/go-bookstore/auth/src/repository/database"
	"github.com/wtran29/go-bookstore/auth/src/repository/rest"
	"github.com/wtran29/go-bookstore/resterr"
)

// type Repository interface {
// 	GetTokenByID(string) (*token.Token, *errors.JsonError)
// 	CreateToken(token.Token) *errors.JsonError
// 	UpdateTokenExpiry(token.Token) *errors.JsonError
// }

type TokenService interface {
	GetTokenByID(string) (*token.Token, *resterr.JsonError)
	CreateToken(token.TokenRequest) (*token.Token, *resterr.JsonError)
	UpdateTokenExpiry(token.Token) *resterr.JsonError
}

type tokenService struct {
	usersRepo rest.UsersRepository
	dbRepo    database.DBRepository
}

func NewService(usersRepo rest.UsersRepository, dbRepo database.DBRepository) TokenService {
	return &tokenService{
		usersRepo: usersRepo,
		dbRepo:    dbRepo,
	}
}

// GetTokenByID implements Service.
func (s *tokenService) GetTokenByID(tokenID string) (*token.Token, *resterr.JsonError) {
	tokenID = strings.TrimSpace(tokenID)
	if len(tokenID) == 0 {
		return nil, resterr.NewBadRequestError("token error", errors.New("invalid access token id"))
	}
	token, err := s.dbRepo.GetTokenByID(tokenID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *tokenService) CreateToken(t token.TokenRequest) (*token.Token, *resterr.JsonError) {
	if err := t.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	user, err := s.usersRepo.LoginUser(t.Username, t.Password)
	if err != nil {
		return nil, err
	}

	at := token.GetNewAccessToken(user.ID)
	at.GenerateToken()

	if err := s.dbRepo.CreateToken(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *tokenService) UpdateTokenExpiry(t token.Token) *resterr.JsonError {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateTokenExpiry(t)
}
