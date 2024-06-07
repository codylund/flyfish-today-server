package user

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSession(c *gin.Context, db *mongo.Database, userID primitive.ObjectID) error {
	// Create a new session and link to user ID.
	sessionsColl := db.Collection("Sessions")
	sessionID := uuid.New().String()
	session := Session{UserID: userID, SessionID: sessionID}
	_, err := sessionsColl.InsertOne(context.TODO(), session)
	if err != nil {
		return err
	}

	// Return secure cookie for the session.
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("session", sessionID, 5*24*60*60, "/", os.Getenv("DOMAIN"), true, true)
	log.Printf("Cookie: %s = %s", "session", sessionID)
	log.Printf("Domain: %s", os.Getenv("DOMAIN"))
	return nil
}
