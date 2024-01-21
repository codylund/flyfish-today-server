package middleware

import (
	"context"
	"log"
	"net/http"
	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func Session(c *gin.Context) {
	sessionId, err := c.Cookie("session")

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Println("SessionID: %s", sessionId)

	db.Run(func (db *mongo.Database) {
		coll := db.Collection("Sessions")
	
		filter := bson.M{"session_id": sessionId}
	
		var result domain.Session
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