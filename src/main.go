package main

import (
	"os"

	"github.com/codylund/streamflows-server/handler"
	"github.com/codylund/streamflows-server/middleware"
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
	unauthenticatedApiGroup.POST("/register", handler.RegisterUser)
	unauthenticatedApiGroup.POST("/signin", handler.SignIn)

	authenticatedApiGroup := router.Group("/v1")
	authenticatedApiGroup.Use(middleware.Session)
	authenticatedApiGroup.GET("/me", handler.Me)
	authenticatedApiGroup.GET("/sites", handler.GetSites)
	authenticatedApiGroup.POST("/signout", handler.SignOut)
	authenticatedApiGroup.POST("/sites/add", handler.AddSite)
	authenticatedApiGroup.PATCH("/sites/:id", handler.UpdateSite)
	authenticatedApiGroup.DELETE("/sites/:id", handler.RemoveSite)

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
