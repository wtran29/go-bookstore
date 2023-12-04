package errors

import (
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	noRowsError = "no rows in result set"
)

func ParseError(err error) *JsonError {
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		switch pgErr.Code {
		case "23505": // violates unique constraint
			return NewBadRequestError("invalid data")
		}
		return NewInternalServerError("error processing request")
	}
	if strings.Contains(err.Error(), noRowsError) {
		return NewNotFoundError("no records matching given id")
	}
	return NewInternalServerError("error parsing database response")

}
