package main

import (
    "github.com/codylund/streamflows-server/handler"
    "github.com/codylund/streamflows-server/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    unauthenticatedGroup := router.Group("/v1")
    unauthenticatedGroup.POST("/register", handler.RegisterUser)
    unauthenticatedGroup.POST("/signin", handler.SignIn)

    authenticatedGroup := router.Group("/v1")
    authenticatedGroup.Use(middleware.Session)
    authenticatedGroup.GET("/sites/get", handler.GetSites)
    authenticatedGroup.POST("/sites/add", handler.AddSites)
    authenticatedGroup.DELETE("/sites/:id", handler.RemoveSite)

    router.Run("localhost:8080")
}
