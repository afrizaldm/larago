package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the user is authenticated
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if user is logged in by checking the token
		token := c.GetHeader("Authorization")

		if token != "valid_token" {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
				"isSuccess": false,
				"message":   "You are not authenticated",
			})
			// If user is not logged in, redirect to the login page
			c.Redirect(http.StatusFound, "/login")
			// Abort the current request
			c.Abort()
		}

		// Continue to the next handler
		c.Next()
	}
}
