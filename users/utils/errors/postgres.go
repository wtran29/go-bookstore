package errors

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	NoRowsError = "no rows in result set"
)

func ParseError(err error) *JsonError {
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		switch pgErr.Code {
		case "23505": // violates unique constraint
			return NewBadRequestError("invalid data")
		case "42601": //syntax error at or near query
			return NewInternalServerError("error occurred in the database")
		}

		return NewInternalServerError(fmt.Sprintf("error processing request %v", err))
	}
	if strings.Contains(err.Error(), NoRowsError) {
		return NewNotFoundError("no records matching given id")
	}
	return NewInternalServerError(fmt.Sprintf("error parsing database response %v", err))

}
