package database

import (
	"errors"

	"github.com/gocql/gocql"
	"github.com/wtran29/go-bookstore/auth/src/clients/cassandra"
	"github.com/wtran29/go-bookstore/auth/src/domain/token"
	"github.com/wtran29/go-bookstore/resterr"
)

const (
	queryGetToken          = "SELECT access_token, user_id, client_id, expiry FROM tokens WHERE access_token=?;"
	queryCreateToken       = "INSERT INTO tokens(access_token, user_id, client_id, expiry) VALUES (?, ?, ?, ?);"
	queryUpdateTokenExpiry = "UPDATE access_tokens SET expiry=? WHERE access_token=?;"
)

type DBRepository interface {
	GetTokenByID(string) (*token.Token, *resterr.JsonError)
	CreateToken(token.Token) *resterr.JsonError
	UpdateTokenExpiry(token.Token) *resterr.JsonError
}

type dbRepository struct {
}

func NewRepo() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetTokenByID(id string) (*token.Token, *resterr.JsonError) {

	var result token.Token
	if err := cassandra.GetSession().Query(queryGetToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expiry,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, resterr.NewNotFoundError("database error", errors.New("no access token found with given id"))
		}

		return nil, resterr.NewInternalServerError("database error", errors.New(err.Error()))
	}

	return &result, nil
}

func (r *dbRepository) CreateToken(t token.Token) *resterr.JsonError {

	if err := cassandra.GetSession().Query(queryCreateToken,
		t.AccessToken,
		t.UserID,
		t.ClientID,
		t.Expiry,
	).Exec(); err != nil {
		return resterr.NewInternalServerError("database error", errors.New(err.Error()))
	}
	return nil
}

func (r *dbRepository) UpdateTokenExpiry(t token.Token) *resterr.JsonError {

	if err := cassandra.GetSession().Query(queryUpdateTokenExpiry,
		t.Expiry,
		t.AccessToken,
	).Exec(); err != nil {
		return resterr.NewInternalServerError("database error", errors.New(err.Error()))
	}
	return nil
}
