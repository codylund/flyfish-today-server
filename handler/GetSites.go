package handler

import (
    "context"
    "github.com/codylund/streamflows-server/db"
    "github.com/codylund/streamflows-server/domain"
    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetSites(c *gin.Context) {
    db.Run(func (db *mongo.Database) {
        coll := db.Collection("Sites")
    
        userID, exists := c.Get("user_id")
        if !exists {
            c.Status(http.StatusBadRequest)
            return
        }


        filter := bson.M{"user_id": userID}
    
        var result domain.SiteCollection
        err := coll.FindOne(context.TODO(), filter).Decode(&result)
        if err != nil {
            //log.Info("Failed to find site collection for user %s: %s", id, err)
            c.Status(http.StatusBadRequest)
            return
        }
        
        c.IndentedJSON(http.StatusOK, result.Sites)
    })
}