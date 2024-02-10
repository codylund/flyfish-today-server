package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddSite(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		Error(c, http.StatusInternalServerError, errors.New("Context missing user id."))
		return
	}

	// Parse sites from request body.
	var site domain.Site
	err := c.BindJSON(&site)
	if err != nil {
		Error(c, http.StatusBadRequest, errors.New("Invalid JSON for new sites."))
		return
	}
	// Add the parsed user id to each site.
	site.UserID = userID.(primitive.ObjectID)

	// TODO verify the sites

	// Insert all the sites.
	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sites")

		// Insert sites.
		insertResult, insertSiteErr := coll.InsertOne(context.TODO(), site)
		if insertSiteErr != nil {
			Error(c, http.StatusInternalServerError, insertSiteErr)
			return
		}

		site.ObjectID = insertResult.InsertedID.(primitive.ObjectID)
		c.IndentedJSON(http.StatusOK, site)
	})
}
