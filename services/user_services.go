package services

import (
	"github.com/jnunortiz/bookstore_users-api/domain/users"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

func CreateUSer(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
