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

func DeleteUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Delete(); err != nil {
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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		err = user.ValidateDataExists()
	} else {
		err = user.ValidatePayload()
	}
	if err != nil {
		return nil, err
	}
	current.UpdateUserInfo(user)
	if err := current.Update(); err != nil {
		return nil, err
	}
	user.DateCreated = current.DateCreated
	return current, nil
}
