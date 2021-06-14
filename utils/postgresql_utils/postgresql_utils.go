package postgresql_utils

import (
	"fmt"
	"strings"

	"github.com/jnunortiz/bookstore_users-api/utils/errors"
	"github.com/lib/pq"
)

const (
	NoRowsError = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), NoRowsError) {
			return errors.NewNotFoundError("no user exists with given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when parsing database error: %s", err.Error()))
	}
	switch sqlErr.Code[0:2] {
	case "23":
		return errors.NewBadRequestError(fmt.Sprintf("%s: %s. %s", sqlErr.Severity, sqlErr.Message, sqlErr.Detail))
	}
	return errors.NewInternalServerError(fmt.Sprintf("%s: %s. %s", sqlErr.Severity, sqlErr.Message, sqlErr.Detail))
}
