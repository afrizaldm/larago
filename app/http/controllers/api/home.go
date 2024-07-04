package api

import (
	"net/http"
	"simple-api/boostrap/module/auth/jwt"

	"github.com/gin-gonic/gin"
)

var secretKey string = "111"

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

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Generate(c *gin.Context) {
	jwtService := jwt.NewJWTService()

	token, _ := jwtService.GenerateToken(secretKey, gin.H{
		"email":      "afrizalmahendra212@gmail.com",
		"username":   "admin",
		"password":   "admin",
		"created_at": 0,
		"updated_at": 0,
		"deleted_at": 0,
	})

	c.String(http.StatusOK, token)
}

func Validate(c *gin.Context) {
	jwtService := jwt.NewJWTService()
	token := c.Param("token")

	data, err := jwtService.ValidateToken(secretKey, token)

	if err == nil {

		c.JSON(http.StatusOK, gin.H{
			"isValid": data,
			"err":     err,
			"data":    data,
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isValid": false,
		"err":     err.Error(),
		"data":    nil,
		"token":   token,
	})

}
