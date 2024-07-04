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

type JwtService struct {
	IJWTService
	issuer  string
	expHour int
}

var jwtService *JwtService = nil

func NewJWTService() *JwtService {

	if jwtService != nil {
		return jwtService
	}

	return &JwtService{
		issuer:  "default_issuer", // Ganti dengan nama penerbit token
		expHour: 24,
	}
}

func (jwtService *JwtService) GenerateToken(secretKey string, value any) (string, error) {
	// Isi dengan kunci rahasia yang aman (sebaiknya dari variabel lingkungan).
	_secretKey := []byte(secretKey)

	// Buat token dengan klaim (misalnya, ID pengguna).
	claims := jwt.MapClaims{
		"value": value,
		"iss":   jwtService.issuer,
		"exp":   time.Now().Add(time.Hour * time.Duration(jwtService.expHour)).Unix(), // Waktu kedaluwarsa.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token menggunakan kunci rahasia.
	tokenString, err := token.SignedString(_secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *JwtService) ValidateToken(secretKey string, tokenString string) (*jwt.Token, error) {
	_secretKey := []byte(secretKey)

	// Parse token dan verifikasi dengan kunci rahasia.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return _secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return token, nil
}
