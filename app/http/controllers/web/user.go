package web

import (
	"net/http"

	"simple-api/app/models"

	"simple-api/app/http/controllers"

	"github.com/gin-gonic/gin"
)

// Struktur untuk request pendaftaran
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

// Struktur untuk response
type RegisterResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struktur untuk response
func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Proses logika bisnis di service
	user := models.User{
		ID:       1, // ini seharusnya didapat dari database
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res := RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	c.JSON(http.StatusOK, res)
}

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
