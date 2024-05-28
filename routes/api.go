package routes

import (
	"net/http"
	"simple-api/controllers"

	"github.com/gin-gonic/gin"
)

func (r *Router) SetupApi() {

	// Daftarkan route dan handler
	r.Engine.POST("/register", controllers.RegisterUser)

	r.Engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
