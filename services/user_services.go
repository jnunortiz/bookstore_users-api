package services

import (
	"github.com/jnunortiz/bookstore_users-api/domain/users"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

func CreateUSer(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
