package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSites(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		Error(c, http.StatusInternalServerError, errors.New("Context missing user id."))
		return
	}

	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sites")

		// Look up site collection for the current user.
		cursor, err := coll.Find(context.TODO(), bson.D{{"user_id", userID}})
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}

		var results []domain.Site
		if err = cursor.All(context.TODO(), &results); err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}
		if results == nil {
			results = []domain.Site{}
		}

		c.IndentedJSON(http.StatusOK, results)
	})
}
