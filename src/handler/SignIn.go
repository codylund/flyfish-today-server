package handler

import (
	"context"
	"errors"
    "github.com/codylund/streamflows-server/auth"
    "github.com/codylund/streamflows-server/db"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
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

        // Create a new session.
        err = auth.NewSession(c, db, userID.ID) 
        if err != nil {
            Error(c, http.StatusInternalServerError, err)
            return
        }
        c.Status(http.StatusOK)
    })
}
