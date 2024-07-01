package web

import (
	"net/http"
	"simple-api/app/models"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard",
	})
}

func Users(c *gin.Context) {

	// Retrieve all users
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "count": result.RowsAffected})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
