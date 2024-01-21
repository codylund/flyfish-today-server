package main

import (
	"context"
	"net/http"
	"github.com/codylund/streamflows-server/auth"
	"github.com/codylund/streamflows-server/db"
	"github.com/codylund/streamflows-server/domain"
	"github.com/codylund/streamflows-server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func register(c *gin.Context) {
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

func signIn(c *gin.Context) {
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

func getSites(c *gin.Context) {
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

func main() {
    router := gin.Default()

	unauthenticatedGroup := router.Group("/v1")
	unauthenticatedGroup.POST("/register", register)
	unauthenticatedGroup.POST("/signin", signIn)

	authenticatedGroup := router.Group("/v1")
	authenticatedGroup.Use(middleware.Session)
	authenticatedGroup.GET("/sites", getSites)

    router.Run("localhost:8080")
}
