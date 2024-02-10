package handler

import (
  "github.com/gin-gonic/gin"
	"net/http"
)

func SignOut(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", true, true)
	c.Status(http.StatusOK)	
}
