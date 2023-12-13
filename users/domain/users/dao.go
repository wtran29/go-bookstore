package users

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/wtran29/go-bookstore/resterr"
	"github.com/wtran29/go-bookstore/users/data/postgres"
	"github.com/wtran29/go-bookstore/users/logger"
	pgerr "github.com/wtran29/go-bookstore/users/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	insertUserStmt    = "INSERT INTO users(first_name, last_name, email, created_at, updated_at, password, status) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;"
	queryGetUser      = "SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE id= $1;"
	insertUpdateUser  = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	insertDeleteUser  = "DELETE FROM users WHERE id=$1;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE status=$1;"
	queryFindByEmail  = "SELECT id, first_name, last_name, email, created_at, updated_at, status, password FROM users WHERE email=$1;"
)

func (user *User) GetUser() *resterr.JsonError {
	query, err := postgres.ClientDB.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error preparing get user query", err)
		return resterr.NewInternalServerError("database error", errors.New("error trying to get user"))
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
		return resterr.NewInternalServerError("database error", errors.New("error trying to get user"))
	}

	return nil
}

func (user *User) SaveUser() *resterr.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertUserStmt)
	if err != nil {
		logger.Error("error preparing save user statement", err)
		return resterr.NewInternalServerError("database error", errors.New("error trying to save user"))
	}
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("error generating hash for password", err)
		return resterr.NewInternalServerError("database error", errors.New("error trying to save user"))
	}

	var userId int64

	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, time.Now(), time.Now(), hashedPassword, user.Status).Scan(&userId)
	if err != nil {
		logger.Error("error getting user id from save user", err)
		return resterr.NewInternalServerError("database error", errors.New("error trying to save user"))
	}

	user.ID = userId

	return nil
}

func (user *User) UpdateUser() *resterr.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertUpdateUser)
	if err != nil {
		logger.Error("error preparing update user statement", err)
		return resterr.NewInternalServerError("database error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error executing update user statement", err)
		return resterr.NewInternalServerError("database error", err)

	}
	return nil
}

func (user *User) DeleteUser() *resterr.JsonError {
	stmt, err := postgres.ClientDB.Prepare(insertDeleteUser)
	if err != nil {
		logger.Error("error preparing delete user statement", err)
		return resterr.NewInternalServerError("database error", err)

	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error executing delete user statement", err)
		return resterr.NewInternalServerError("database error", err)

	}
	return nil
}

func (user *User) FindUserByStatus(status string) ([]User, *resterr.JsonError) {
	query, err := postgres.ClientDB.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error preparing find user by status query", err)
		return nil, resterr.NewInternalServerError("database error", errors.New("error when trying to find user"))

	}
	defer query.Close()

	rows, err := query.Query(status)
	if err != nil {
		logger.Error("error query find user by status rows", err)
		return nil, resterr.NewInternalServerError("database error", errors.New("error when trying to find user"))
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status); err != nil {
			logger.Error("error looping through find user by status rows", err)
			return nil, resterr.NewInternalServerError("database error", errors.New("error when trying to find user"))
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, resterr.NewNotFoundError("database error", errors.New(fmt.Sprintf("no users matching status %s", status)))
	}
	return users, nil
}

func (user *User) FindByEmail() (*User, *resterr.JsonError) {
	query, err := postgres.ClientDB.Prepare(queryFindByEmail)
	if err != nil {
		logger.Error("error preparing find user by email query", err)
		return nil, resterr.NewInternalServerError("database error", errors.New("error when trying to find email"))
	}
	defer query.Close()

	row := query.QueryRow(user.Email)

	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Status,
		&user.Password,
	); err != nil {
		if strings.Contains(err.Error(), pgerr.NoRowsError) {
			return nil, resterr.NewNotFoundError("database error", errors.New("user not found"))
		}
		logger.Error("error retrieving user by email", err)
		return nil, resterr.NewInternalServerError("database error", errors.New("error when trying to find email"))
	}

	fmt.Printf("User from dao.go:%v\n", user)

	return user, nil
}
