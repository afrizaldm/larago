package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard",
	})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
