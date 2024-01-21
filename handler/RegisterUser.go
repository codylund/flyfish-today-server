package handler

import (
	"context"
    "github.com/codylund/streamflows-server/auth"
    "github.com/codylund/streamflows-server/db"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    // Parse user information from request body.
    var reqBody domain.User
    err := c.Bind(&reqBody)
    if err != nil {
        c.Status(http.StatusBadRequest)
        return
    }

    db.Run(func (db *mongo.Database) {
        coll := db.Collection("Users")
        
        // Verify no existing user with same username.
        count, err := coll.CountDocuments(context.TODO(), bson.M{"username": reqBody.Username})
        if err != nil || count > 0 {
            c.Status(http.StatusBadRequest)
            return
        }

        // Create new user.
        reqBody.Password, _ = auth.HashPassword(reqBody.Password)
        _, createUserErr := coll.InsertOne(context.TODO(), reqBody)
        if createUserErr != nil {
            c.Status(http.StatusBadRequest)
            return
        }

        c.Status(http.StatusOK)	
    })
}