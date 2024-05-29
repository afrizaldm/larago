package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the dashboard",
	})
}

func Hallo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hallo",
	})
}
