package api

import (
	"net/http"

	"simple-api/models"

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
