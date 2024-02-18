package user

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignIn(c *gin.Context) {
	userRequest, err := GetUser(c)
	if err != nil {
		util.Error(c, http.StatusBadRequest, err)
		return
	}

	db.Run(func(db *mongo.Database) {
		usersColl := db.Collection("Users")

		// Lookup by username.
		result := usersColl.FindOne(
			context.TODO(),
			bson.M{"username": strings.ToLower(userRequest.Username)},
		)

		// Decode password hash from DB.
		var user User
		err = result.Decode(&user)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}

		// Verify password hash.
		if !CheckPasswordHash(userRequest.Password, user.Password) {
			util.Error(c, http.StatusUnauthorized, errors.New(user.Password))
			return
		}

		// Password matched! Decode user ID from DB.
		var userID UserID
		err = result.Decode(&userID)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}

		// Create a new session.
		err = NewSession(c, db, userID.ID)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)
	})
}
