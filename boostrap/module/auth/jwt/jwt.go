package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService menyediakan fungsi untuk membuat dan memvalidasi token JWT
type IJWTService interface {
	GenerateToken(username string, isAdmin bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

// type jwtCustomClaims struct {
// 	Username string `json:"username"`
// 	IsAdmin  bool   `json:"is_admin"`
// }

type jwtService struct {
	secretKey string
	issuer    string
	expHour   int
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret_key",  // Ganti dengan kunci rahasia yang lebih kuat
		issuer:    "issuer_name", // Ganti dengan nama penerbit token
		expHour:   24,
	}
}

func (jwtService *jwtService) GenerateToken(value any) (string, error) {
	// Isi dengan kunci rahasia yang aman (sebaiknya dari variabel lingkungan).
	secretKey := []byte(jwtService.secretKey)

	// Buat token dengan klaim (misalnya, ID pengguna).
	claims := jwt.MapClaims{
		"value": "",
		"iss":   jwtService.issuer,
		"exp":   time.Now().Add(time.Hour * time.Duration(jwtService.expHour)).Unix(), // Waktu kedaluwarsa.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token menggunakan kunci rahasia.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(jwtService.secretKey)

	// Parse token dan verifikasi dengan kunci rahasia.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	return token, nil
}
