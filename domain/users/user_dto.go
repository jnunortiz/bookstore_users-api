package users

import (
	"strings"

	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) ValidateFirstName() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid or missing first name")
	}
	return nil
}

func (user *User) ValidateLastName() *errors.RestErr {
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid or missing last name")
	}
	return nil
}

func (user *User) ValidateEmail() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid or missing email address")
	}
	return nil
}

func (user *User) ValidatePayload() *errors.RestErr {
	if err := user.ValidateFirstName(); err != nil {
		return err
	}
	if err := user.ValidateLastName(); err != nil {
		return err
	}
	if err := user.ValidateEmail(); err != nil {
		return err
	}
	return nil
}
