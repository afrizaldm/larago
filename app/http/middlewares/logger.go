package middlewares

import (
	"fmt"
	"simple-api/boostrap/module/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs details about each request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log details
		duration := time.Since(start)
		status := c.Writer.Status()
		logValue := fmt.Sprintf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, status, duration)

		logger.Write(logValue)
	}
}
