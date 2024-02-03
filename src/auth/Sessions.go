package auth

import (
	"context"
	"github.com/codylund/streamflows-server/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func Error(c *gin.Context, status int, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(status , gin.H{"message": err.Error()})
}

func NewSession(c *gin.Context, db *mongo.Database, userID primitive.ObjectID) error {
	// Create a new session and link to user ID.
	sessionsColl := db.Collection("Sessions")
	sessionID := uuid.New().String()
	session := domain.Session{UserID: userID, SessionID: sessionID}
	_, err := sessionsColl.InsertOne(context.TODO(), session)
	if err != nil {
			return err
	}

	// Return secure cookie for the session.
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("session", sessionID, 180*24*60*60, "/", os.Getenv("SERVER_ADDRESS"), true, true)
	return nil
}
