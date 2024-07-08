package web

import (
	"net/http"
	"simple-api/app/models"

	"github.com/gin-contrib/sessions"
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

func AuthLogin(c *gin.Context) {

	email := c.Query("email")
	password := c.Query("password")

	var user models.User

	models.DB.First(&user, "email = ?", email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isSuccess": false,
			"message":   "Username not found",
		})

		return
	}

	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isSuccess": false,
			"message":   "Password not match",
		})

		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	c.JSON(200, gin.H{
		"isSuccess": true,
		"message":   "Login successful",
	})

}

func AuthLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{
		"isSuccess": true,
		"message":   "Logout successful",
	})
}

func AuthCheck(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")

	if user == nil {
		c.JSON(401, gin.H{
			"isSuccess": false,
			"message":   "Unauthorized",
			"user":      nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"isSuccess": true,
		"message":   "Welcome " + user.(string),
		"user":      user,
	})

}

func AuthPing(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")

	if user == nil {
		c.String(http.StatusUnauthorized, "")
		return
	}
	c.String(http.StatusOK, "pong")
}
