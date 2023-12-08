package database

import (
	"github.com/gocql/gocql"
	"github.com/wtran29/go-bookstore/auth/src/clients/cassandra"
	"github.com/wtran29/go-bookstore/auth/src/domain/token"

	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
)

const (
	queryGetToken          = "SELECT access_token, user_id, client_id, expiry FROM tokens WHERE access_token=?;"
	queryCreateToken       = "INSERT INTO tokens(access_token, user_id, client_id, expiry) VALUES (?, ?, ?, ?);"
	queryUpdateTokenExpiry = "UPDATE access_tokens SET expiry=? WHERE access_token=?;"
)

type DBRepository interface {
	GetTokenByID(string) (*token.Token, *errors.JsonError)
	CreateToken(token.Token) *errors.JsonError
	UpdateTokenExpiry(token.Token) *errors.JsonError
}

type dbRepository struct {
}

func NewRepo() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetTokenByID(id string) (*token.Token, *errors.JsonError) {

	var result token.Token
	if err := cassandra.GetSession().Query(queryGetToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expiry,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}

		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) CreateToken(t token.Token) *errors.JsonError {

	if err := cassandra.GetSession().Query(queryCreateToken,
		t.AccessToken,
		t.UserID,
		t.ClientID,
		t.Expiry,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateTokenExpiry(t token.Token) *errors.JsonError {

	if err := cassandra.GetSession().Query(queryUpdateTokenExpiry,
		t.Expiry,
		t.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
