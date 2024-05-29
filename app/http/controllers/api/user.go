package api

import (
	"net/http"
	"simple-api/app/http/controllers"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.ControllerResources
}

func (ctr *UserController) Index(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Index",
	})
}

func (ctr *UserController) Create(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Create",
	})
}

func (ctr *UserController) Store(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Store",
	})
}

func (ctr *UserController) Show(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Show",
	})
}

func (ctr *UserController) Edit(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit",
	})
}

func (ctr *UserController) Update(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Update",
	})
}

func (ctr *UserController) Destroy(c *gin.Context) {
	// Your custom logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Destroy",
	})
}
