package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs details about each request
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log details
		duration := time.Since(start)
		status := c.Writer.Status()
		fmt.Printf("%s %s %d %s\n", c.Request.Method, c.Request.URL.Path, status, duration)
	}
}

// AuthMiddleware checks if the user is authenticated
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if user is logged in by checking the token
		token := c.GetHeader("Authorization")

		if token != "valid_token" { // Replace with your actual token validation
			// If user is not logged in, redirect to the login page
			c.Redirect(http.StatusFound, "/login")
			// Abort the current request
			c.Abort()
		}

		// Continue to the next handler
		c.Next()
	}
}
