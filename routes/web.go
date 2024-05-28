package routes

import (
	"net/http"
	"simple-api/controllers"

	"github.com/gin-gonic/gin"
)

type WebRouter struct {
	Engine *gin.Engine
}

func (router *WebRouter) SetupRouter() *WebRouter {
	r := gin.Default()

	// Daftarkan route dan handler
	r.POST("/register", controllers.RegisterUser)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Engine = r

	return router
}

func (router *WebRouter) GetEngine() *gin.Engine {
	return router.Engine
}
