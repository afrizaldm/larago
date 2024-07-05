package api

import (
	"net/http"
	"simple-api/boostrap/module/auth/jwt"
	"simple-api/config"

	"github.com/gin-gonic/gin"
)

// global data in api
var appConfig *config.IAppConfig = config.NewAppConfig().Load()

func AuthLogin(c *gin.Context) {
	jwtService := jwt.NewJWTService()

	token, _ := jwtService.GenerateToken(appConfig.APP_SECRET_KEY, gin.H{
		"email":      "afrizalmahendra212@gmail.com",
		"username":   "admin",
		"password":   "admin",
		"created_at": 0,
		"updated_at": 0,
		"deleted_at": 0,
	})

	c.String(http.StatusOK, token)
}

func AuthLogout(c *gin.Context) {

}

func AuthUser(c *gin.Context) {

	jwtService := jwt.NewJWTService()
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

func RefreshToken(c *gin.Context) {

}
