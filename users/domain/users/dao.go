package users

import (
	"fmt"
	"time"

	"github.com/wtran29/go-bookstore/users/data/postgres"
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
		&user.Status,
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	var userId int64

	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, time.Now(), time.Now(), hash, user.Status).Scan(&userId)
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

func (user *User) FindUserByStatus(status string) ([]User, *errors.JsonError) {
	query, err := postgres.ClientDB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer query.Close()

	rows, err := query.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status); err != nil {
			return nil, errors.ParseError(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return users, nil
}
