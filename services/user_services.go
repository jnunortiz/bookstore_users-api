package services

import (
	"github.com/jnunortiz/bookstore_users-api/domain/users"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUSer(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidatePayload(); err != nil {
		return nil, err
	}
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := current.Update(user); err != nil {
		return nil, err
	}
	user.DateCreated = current.DateCreated
	return &user, nil
}
