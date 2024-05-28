package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.templ", gin.H{
		"title": "Dashboard",
	})
}
