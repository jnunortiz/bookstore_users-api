package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnunortiz/bookstore_users-api/domain/users"
	"github.com/jnunortiz/bookstore_users-api/services"
	"github.com/jnunortiz/bookstore_users-api/utils/errors"
)

func CreateUSer(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUSer(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
