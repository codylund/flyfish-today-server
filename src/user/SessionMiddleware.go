package user

import (
	"context"
	"net/http"

	"github.com/codylund/streamflows-server/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SessionMiddleware(c *gin.Context) {
	sessionId, err := c.Cookie("session")

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sessions")

		filter := bson.M{"session_id": sessionId}

		var result Session
		err := coll.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			//log.Info("Failed to find site collection for user %s: %s", id, err)
			c.Status(http.StatusBadRequest)
			return
		}

		// Set user property.
		c.Set("user_id", result.UserID)

		c.Next()
	})
}
