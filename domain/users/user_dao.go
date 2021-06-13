package users

import (
	"fmt"
	"strings"

	"github.com/jnunortiz/bookstore_users-api/datasources/postgres/users_db"
	"github.com/jnunortiz/bookstore_users-api/utils/date_utils"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id;"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = $1;"
	UniqueEmailError = "users_email_key"
	NoRowsError      = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		error := err.Error()
		if strings.Contains(error, NoRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("user with id %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when getting user id %d: %s", user.Id, error))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&user.Id)
	if err != nil {
		error := err.Error()
		if strings.Contains(error, UniqueEmailError) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to insert user. %s", error))
	}
	return nil
}
