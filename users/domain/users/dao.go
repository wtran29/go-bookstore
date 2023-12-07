package users

import (
	"fmt"
	"time"

	"github.com/wtran29/go-bookstore/users/data/postgres"
	"github.com/wtran29/go-bookstore/users/logger"
	"github.com/wtran29/go-bookstore/users/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	insertUserStmt        = "INSERT INTO users(first_name, last_name, email, created_at, updated_at, password, status) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;"
	queryGetUser          = "SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE id= $1;"
	insertUpdateUser      = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	insertDeleteUser      = "DELETE FROM users WHERE id=$1;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE status=$1;"
)

func (user *User) GetUser() *errors.JsonError {
	query, err := postgres.ClientDB.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error preparing get user query", err)
		return errors.ParseError(err)
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
		&user.Status,
	); err != nil {
		logger.Error("error retrieving id from get user", err)
		return errors.ParseError(err)
	}

	return nil
}

func (user *User) SaveUser() *errors.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertUserStmt)
	if err != nil {
		logger.Error("error preparing save user statement", err)
		return errors.ParseError(err)
	}
	defer stmt.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("error generating hash for password", err)
		return errors.NewInternalServerError(err.Error())
	}

	var userId int64

	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, time.Now(), time.Now(), hash, user.Status).Scan(&userId)
	if err != nil {
		logger.Error("error getting user id from save user", err)
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
		logger.Error("error preparing update user statement", err)
		return errors.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error executing update user statement", err)
		return errors.ParseError(err)
	}
	return nil
}

func (user *User) DeleteUser() *errors.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertDeleteUser)
	if err != nil {
		logger.Error("error preparing delete user statement", err)
		return errors.ParseError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error executing delete user statement", err)
		return errors.ParseError(err)
	}
	return nil
}

func (user *User) FindUserByStatus(status string) ([]User, *errors.JsonError) {
	query, err := postgres.ClientDB.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error preparing find user by status query", err)
		return nil, errors.ParseError(err)
	}
	defer query.Close()

	rows, err := query.Query(status)
	if err != nil {
		logger.Error("error query find user by status rows", err)
		return nil, errors.ParseError(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status); err != nil {
			logger.Error("error looping through find user by status rows", err)
			return nil, errors.ParseError(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return users, nil
}
