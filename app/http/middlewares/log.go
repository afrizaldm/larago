package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs details about each request
func Log() gin.HandlerFunc {
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
