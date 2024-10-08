package sites

import (
	"context"
	"errors"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateSite(c *gin.Context) {
	// Get session user id from the current Gin context.
	// This is set by middleware.Sessions.
	userID, exists := c.Get("user_id")
	if !exists {
		util.Error(c, http.StatusInternalServerError, errors.New("Context missing user id."))
		return
	}

	siteID, siteIDError := primitive.ObjectIDFromHex(c.Param("id"))
	if siteIDError != nil {
		util.Error(c, http.StatusBadRequest, siteIDError)
		return
	}

	filter := bson.M{"_id": siteID, "user_id": userID}

	var siteUpdate SiteUpdate
	bindJsonErr := c.BindJSON(&siteUpdate)
	if bindJsonErr != nil {
		util.Error(c, http.StatusBadRequest, bindJsonErr)
		return
	}

	// siteUpdateBson, marshalError := json.Marshal(siteUpdate)
	// if marshalError != nil {
	// 	Error(c, http.StatusBadRequest, marshalError)
	// 	return
	// }

	update := bson.D{{Key: "$set", Value: siteUpdate}}

	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sites")

		// Look up site collection for the current user.
		_, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	})
}
