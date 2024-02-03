package main

import (
    "github.com/codylund/streamflows-server/handler"
    "github.com/codylund/streamflows-server/middleware"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "os"
)

func main() {
    router := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://0.0.0.0:3000"}
    config.AllowCredentials = true
    router.Use(cors.New(config))

    unauthenticatedGroup := router.Group("/v1")
    unauthenticatedGroup.POST("/register", handler.RegisterUser)
    unauthenticatedGroup.POST("/signin", handler.SignIn)

    authenticatedGroup := router.Group("/v1")
    authenticatedGroup.Use(middleware.Session)
    authenticatedGroup.GET("/me", handler.Me)
    authenticatedGroup.GET("/sites/get", handler.GetSites)
    authenticatedGroup.POST("/signout", handler.SignOut)
    authenticatedGroup.POST("/sites/add", handler.AddSite)
    authenticatedGroup.DELETE("/sites/:id", handler.RemoveSite)

    router.Run(os.Getenv("SERVER_ADDRESS"))
}
