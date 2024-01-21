package handler

import (
	"context"
	"errors"
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
    userRequest, err := GetUser(c)
    if err != nil {
        Error(c, http.StatusBadRequest, err)
        return
    }

    db.Run(func (db *mongo.Database) {
        usersColl := db.Collection("Users")
        
        // Lookup by username.
        result := usersColl.FindOne(context.TODO(), bson.M{"username": userRequest.Username})

        // Decode password hash from DB.
        var user domain.User
        err = result.Decode(&user)
        if err != nil {
			Error(c, http.StatusInternalServerError, err)
            return
        }
        
        // Verify password hash.
        if !auth.CheckPasswordHash(userRequest.Password, user.Password) {
			Error(c, http.StatusUnauthorized, errors.New(user.Password))
            return
        }

        // Password matched! Decode user ID from DB.
        var userID domain.UserID
        err = result.Decode(&userID)
        if err != nil {
            Error(c, http.StatusInternalServerError, err)
            return
        }

        // Create a new session and link to user ID.
        sessionsColl := db.Collection("Sessions")
		sessionID := uuid.New().String()
		session := domain.Session{UserID: userID.ID, SessionID: sessionID}
        _, err = sessionsColl.InsertOne(context.TODO(), session)
        if err != nil {
            Error(c, http.StatusInternalServerError, err)
            return
        }

		// Return secure cookie for the session.
        c.SetCookie("session", sessionID, 180*24*60*60, "/", "localhost", true, true)
        c.Status(http.StatusOK)	
    })
}