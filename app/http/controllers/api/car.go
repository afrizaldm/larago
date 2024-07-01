package api

import (
	"net/http"
	"simple-api/app/http/controllers"
	"simple-api/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	controllers.ControllerResources
}

func (ctr *CarController) Index(c *gin.Context) {

	// Retrieve all cars
	var cars []models.Car
	result := models.DB.Find(&cars)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars, "count": result.RowsAffected})
}

func (ctr *CarController) Create(c *gin.Context) {

	// Return Form Data
	c.JSON(500, gin.H{"error": "Failed to retrieve car"})
}

func (ctr *CarController) Store(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Create
	models.DB.Create(&models.Car{ID: id, Name: name, Email: email, Password: password})
}

func (ctr *CarController) Show(c *gin.Context) {

	id := c.Param("car")

	// Get Car by ID
	var car models.Car
	result := models.DB.First(&car, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car, "count": result.RowsAffected})
}

func (ctr *CarController) Edit(c *gin.Context) {

	id := c.Param("car")

	// Get Car by ID
	var car models.Car
	result := models.DB.First(&car, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car, "count": result.RowsAffected})
}

func (ctr *CarController) Update(c *gin.Context) {

	id := c.Param("car")

	// Get Car by ID
	var car models.Car
	result := models.DB.First(&car, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve car"})
		return
	}

	// Update - update product's price to 200
	models.DB.Model(&car).Update("name", 200)

	// Update - update multiple fields
	models.DB.Model(&car).Updates(models.Car{Name: "200", Email: "F42"}) // non-zero fields
	models.DB.Model(&car).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func (ctr *CarController) Destroy(c *gin.Context) {

	// Delete - delete product
	var car models.Car
	models.DB.Delete(&car, 1)
}
