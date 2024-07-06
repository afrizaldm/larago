package api

import (
	"net/http"
	"simple-api/boostrap/module/auth/jwt"

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

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Generate(c *gin.Context) {
	jwtService := jwt.Instance()

	tokens, err := jwtService.GenerateTokens(gin.H{
		"email":      "afrizalmahendra212@gmail.com",
		"username":   "admin",
		"password":   "admin",
		"created_at": 0,
		"updated_at": 0,
		"deleted_at": 0,
	}, appConfig.APP_SECRET_KEY, appConfig.APP_SECRET_KEY_REFRESH_TOKEN)

	// c.String(http.StatusOK, "pong")
	// return

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"tokens": tokens,
			"err":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
		"err":    err,
		"mess":   "must error",
	})
}

func Validate(c *gin.Context) {
	jwtService := jwt.Instance()
	token := c.Param("token")

	data, err := jwtService.ValidateToken(appConfig.APP_SECRET_KEY, token)

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
