package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// csrfMiddleware := csrf.Protect(
		// 	[]byte("32-byte-long-auth-key"),
		// 	csrf.Secure(false), // Set true di production
		// )

		// csrfMiddleware(c.Handler())
		c.Next()
	}
}
