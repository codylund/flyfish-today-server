package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", true, true)
	c.Status(http.StatusOK)
}
