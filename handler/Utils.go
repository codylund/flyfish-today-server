package handler

import (
    "errors"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (domain.User, error) {
    var user domain.User
    err := c.BindJSON(&user)
    if err != nil {
        return user, err
    }

    if user.Username == "" || user.Password == "" {
        return user, errors.New("Invalid user input.")
    }

    return user, err
}

func Error(c *gin.Context, status int, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(status , gin.H{"message": err.Error()})
}