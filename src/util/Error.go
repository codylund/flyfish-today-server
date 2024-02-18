package util

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, status int, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(status, gin.H{"message": err.Error()})
}
