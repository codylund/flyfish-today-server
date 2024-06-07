package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Me(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		util.Error(c, http.StatusInternalServerError, errors.New("context missing user id"))
		return
	}

	db.Run(func(db *mongo.Database) {
		usersColl := db.Collection("Users")

		// Lookup by username.
		result := usersColl.FindOne(context.TODO(), bson.M{"_id": userID})

		// Decode password hash from DB.
		var profile Profile
		err := result.Decode(&profile)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, profile)
	})
}
