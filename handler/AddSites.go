package handler

import (
	"context"
    "errors"
    "github.com/codylund/streamflows-server/db"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
)

func AddSites(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		Error(c, http.StatusInternalServerError, errors.New("Context missing user id."))
		return
	}

	// Parse sites from request body.
	var sitesArray []domain.Site
    err := c.BindJSON(&sitesArray)
    if err != nil {
        Error(c, http.StatusBadRequest, errors.New("Invalid JSON for new sites."))
		return
    }
	// Add the parsed user id to each site.
	var sites []interface{}
	for _, site := range sitesArray {
		site.UserID = userID.(primitive.ObjectID)
		sites = append(sites, site)
	}

	// TODO verify the sites

	// Insert all the sites.
    db.Run(func (db *mongo.Database) {
        coll := db.Collection("Sites")

		// Insert sites.
        _, insertSitesErr := coll.InsertMany(context.TODO(), sites)
        if insertSitesErr != nil {
            Error(c, http.StatusInternalServerError, insertSitesErr)
            return
        }
        
        c.Status(http.StatusOK)	
    })
}