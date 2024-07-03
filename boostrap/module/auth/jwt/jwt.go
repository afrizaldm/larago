package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService menyediakan fungsi untuk membuat dan memvalidasi token JWT
type JWTService interface {
	GenerateToken(username string, isAdmin bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "secret_key",  // Ganti dengan kunci rahasia yang lebih kuat
		issuer:    "issuer_name", // Ganti dengan nama penerbit token
	}
}

func (j *jwtService) GenerateToken(username string, isAdmin bool) (string, error) {
	claims := &jwtCustomClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
