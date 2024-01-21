package handler

import (
	"context"
    "github.com/codylund/streamflows-server/auth"
    "github.com/codylund/streamflows-server/db"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SignIn(c *gin.Context) {
    // Parse user information from request body.
    var reqBody domain.User
    err := c.Bind(&reqBody)
    if err != nil {
        c.Status(http.StatusBadRequest)
        return
    }

    db.Run(func (db *mongo.Database) {
        usersColl := db.Collection("Users")
        
        // Lookup by username.
        result := usersColl.FindOne(context.TODO(), bson.M{"username": reqBody.Username})

        // Decode password hash from DB.
        var user domain.User
        decodeUserErr := result.Decode(&user)
        if decodeUserErr != nil {
            c.Status(http.StatusBadRequest)
            return
        }
        
        // Verify password hash.
        if !auth.CheckPasswordHash(reqBody.Password, user.Password) {
            c.Status(http.StatusUnauthorized)
            return
        }

        // Password matched! Decode user ID from DB.
        var userID domain.UserID
        decodeUserIDErr := result.Decode(&user)
        if decodeUserIDErr != nil {
            c.Status(http.StatusBadRequest)
            return
        }

        // Create a new session.
        sessionsColl := db.Collection("Users")
        sessionID := uuid.New().String()
        _, createSessionErr := sessionsColl.InsertOne(context.TODO(), domain.Session{UserID: userID.ID, SessionID: sessionID})
        if createSessionErr != nil {
            c.Status(http.StatusBadRequest)
            return
        }

        c.SetCookie("session", sessionID, 180*24*60*60, "/", "localhost", true, true)
        c.Status(http.StatusOK)	
    })
}