package main

import (
	"os"

	"github.com/codylund/streamflows-server/sites"
	"github.com/codylund/streamflows-server/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	origin := os.Getenv("ORIGIN_URL")
	println("Allowing origin: ", origin)
	config.AllowOrigins = []string{origin}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	unauthenticatedApiGroup := router.Group("/v1")
	unauthenticatedApiGroup.POST("/register", user.RegisterUser)
	unauthenticatedApiGroup.POST("/signin", user.SignIn)

	authenticatedApiGroup := router.Group("/v1")
	authenticatedApiGroup.Use(user.SessionMiddleware)

	authenticatedApiGroup.GET("/me", user.Me)
	authenticatedApiGroup.POST("/signout", user.SignOut)

	authenticatedApiGroup.GET("/sites", sites.GetSites)
	authenticatedApiGroup.POST("/sites/add", sites.AddSite)
	authenticatedApiGroup.PATCH("/sites/:id", sites.UpdateSite)
	authenticatedApiGroup.DELETE("/sites/:id", sites.RemoveSite)

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
