package middlewares

import (
	"net/http"
	"simple-api/boostrap/module/auth/jwt"
	"simple-api/config"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the user is authenticated using jwt
func IsAuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil nilai header Authorization
		authHeader := c.GetHeader("Authorization")

		// Memeriksa apakah Authorization header kosong atau tidak
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Authorization header is required."})
			c.Abort()
			return
		}

		// Memisahkan nilai Authorization header untuk mendapatkan token
		// Format: "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Malformed authorization header."})
			c.Abort()
			return
		}

		jwtService := jwt.NewJWTService()
		appConfig := config.NewAppConfig().Load()

		// Mengambil token dari array yang telah dipisahkan
		token := bearerToken[1]

		// Lakukan sesuatu dengan token, misalnya verifikasi atau pemrosesan lebih lanjut
		// Di sini kamu dapat memanggil fungsi untuk memverifikasi token JWT, misalnya
		// validateToken(token)

		// check if token valid
		data, err := jwtService.ValidateToken(appConfig.APP_SECRET_KEY, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Unauthorized: Invalid token."})
			c.Abort()
			return
		}

		// Setel data token ke dalam konteks Gin untuk digunakan di handler selanjutnya jika perlu
		c.Set("token", token)
		c.Set("auth-data", data)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
