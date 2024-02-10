package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RemoveSite(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		Error(c, http.StatusInternalServerError, errors.New("Context missing user id."))
		return
	}

	siteID, siteIDError := primitive.ObjectIDFromHex(c.Param("id"))
	if siteIDError != nil {
		Error(c, http.StatusBadRequest, siteIDError)
		return
	}

	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sites")

		// Look up site collection for the current user.
		result, err := coll.DeleteMany(context.TODO(), bson.M{"_id": siteID, "user_id": userID})
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
			return
		}

		if result.DeletedCount <= 0 {
			Error(c, http.StatusBadRequest, errors.New("Site not found."))
			return
		}
		c.Status(http.StatusOK)
	})
}
