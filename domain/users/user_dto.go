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

func (user *User) trimUser() {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
}

func (user *User) ValidateDataExists() *errors.RestErr {
	user.trimUser()
	if user.Email == "" && user.FirstName == "" && user.LastName == "" {
		return errors.NewBadRequestError("no information received for update")
	}
	return nil
}

func (user *User) ValidateEmail() *errors.RestErr {
	user.trimUser()
	if user.Email == "" {
		return errors.NewBadRequestError("invalid or missing email address")
	}
	return nil
}

func (user *User) ValidatePayload() *errors.RestErr {
	if err := user.ValidateEmail(); err != nil {
		return err
	}
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid or missing first name")
	}
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid or missing last name")
	}
	return nil
}

func (user *User) UpdateUserInfo(NewUser User) {
	if NewUser.FirstName != "" {
		user.FirstName = NewUser.FirstName
	}
	if NewUser.LastName != "" {
		user.LastName = NewUser.LastName
	}
	if NewUser.Email != "" {
		user.Email = NewUser.Email
	}
}
