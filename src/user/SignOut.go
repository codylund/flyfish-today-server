package user

import (
	"context"
	"net/http"
	"os"

	"github.com/codylund/streamflows-server/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignOut(c *gin.Context) {
	// Make sure there
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.Status(http.StatusOK)
		return
	}

	// Clean up the session in the DB.
	db.Run(func(db *mongo.Database) {
		coll := db.Collection("Sessions")
		filter := bson.M{"session_id": sessionId}
		_, err := coll.DeleteMany(context.TODO(), filter)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	})

	// Clear out the cookie.
	c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
	c.Status(http.StatusOK)
}
