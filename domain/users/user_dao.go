package users

import (
	"github.com/jnunortiz/bookstore_users-api/datasources/postgres/users_db"
	"github.com/jnunortiz/bookstore_users-api/utils/date_utils"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
	"github.com/jnunortiz/bookstore_users-api/utils/postgresql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id;"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = $1;"
	queryUpdateUser = "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	NoRowsError     = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return postgresql_utils.ParseError(err)
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
		return postgresql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return postgresql_utils.ParseError(err)
	}
	return nil
}
