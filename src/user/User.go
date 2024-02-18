package user

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username    string `json:"username"     bson:"username"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Password    string `json:"password"     bson:"password"`
}

func GetUser(c *gin.Context) (User, error) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		return user, err
	}

	if user.Username == "" || user.Password == "" {
		return user, errors.New("Invalid user input.")
	}

	return user, err
}
