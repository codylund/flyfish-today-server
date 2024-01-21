package handler

import (
	"context"
    "errors"
    "fmt"
    "github.com/codylund/streamflows-server/auth"
    "github.com/codylund/streamflows-server/db"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    user, err := GetUser(c)
    if err != nil {
        Error(c, http.StatusBadRequest, err)
        return
    }

    db.Run(func (db *mongo.Database) {
        coll := db.Collection("Users")
        
        // Verify no existing user with same username.
        count, err := coll.CountDocuments(context.TODO(), bson.M{"username": user.Username})
        if err != nil {
            Error(c, http.StatusInternalServerError, err)
            return
        }
        if count > 0 {
            msg := fmt.Sprintf("An account with username %s already exists.", user.Username)
            Error(c, http.StatusBadRequest, errors.New(msg))
            return
        }

        // Create new user with hashed password.
        user.Password, _ = auth.HashPassword(user.Password)
        _, createUserErr := coll.InsertOne(context.TODO(), user)
        if createUserErr != nil {
            Error(c, http.StatusInternalServerError, createUserErr)
            return
        }
        
        c.Status(http.StatusOK)	
    })
}