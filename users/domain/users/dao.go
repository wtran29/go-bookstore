package users

import (
	"time"

	"github.com/wtran29/go-bookstore/users/data/postgres"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

const (
	insertUserStmt   = "INSERT INTO users(first_name, last_name, email, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	queryGetUser     = "SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id= $1;"
	insertUpdateUser = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	insertDeleteUser = "DELETE FROM users WHERE id=$1;"
)

func (user *User) GetUser() *errors.JsonError {
	query, err := postgres.ClientDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer query.Close()

	row := query.QueryRow(user.ID)

	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return errors.ParseError(err)
	}

	return nil
}

func (user *User) SaveUser() *errors.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertUserStmt)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var userId int64

	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, time.Now(), time.Now()).Scan(&userId)
	if err != nil {
		return errors.ParseError(err)
	}

	user.ID = userId

	// user.CreatedAt = date.GetNow()
	// usersDB[user.ID] = user
	return nil
}

func (user *User) UpdateUser() *errors.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.ParseError(err)
	}
	return nil
}

func (user *User) DeleteUser() *errors.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return errors.ParseError(err)
	}
	return nil
}
