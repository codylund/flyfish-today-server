package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUser(c *gin.Context) {
	user, err := GetUser(c)
	if err != nil {
		util.Error(c, http.StatusBadRequest, err)
		return
	}

	// Sanitize the user input.
	user.Username = strings.ToLower(user.Username)

	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Users")

		// Verify no existing user with same username.
		count, err := coll.CountDocuments(context.TODO(), bson.M{"username": user.Username})
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		if count > 0 {
			msg := fmt.Sprintf("An account with username %s already exists.", user.Username)
			util.Error(c, http.StatusBadRequest, errors.New(msg))
			return
		}

		// Create new user with hashed password.
		user.Password, _ = HashPassword(user.Password)
		res, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}

		// Start a new session for the user.
		_ = NewSession(c, db, res.InsertedID.(primitive.ObjectID))
		c.Status(http.StatusOK)
	})
}
