package errors

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wtran29/go-bookstore/resterr"
)

const (
	NoRowsError = "no rows in result set"
)

func ParseError(err error) *resterr.JsonError {
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		switch pgErr.Code {
		case "23505": // violates unique constraint
			return resterr.NewBadRequestError("invalid data", err)
		case "42601": //syntax error at or near query
			return resterr.NewInternalServerError("error occurred in the database", err)
		}

		return resterr.NewInternalServerError("error processing request", err)
	}
	if strings.Contains(err.Error(), NoRowsError) {
		return resterr.NewNotFoundError("no records matching given id", err)
	}
	return resterr.NewInternalServerError("error parsing database response", errors.New("database error"))

}
