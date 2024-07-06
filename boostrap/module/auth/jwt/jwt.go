package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
}

type Claims struct {
	Value any `json:"username"`
	jwt.RegisteredClaims
}

var jwtService *JWTService = nil

func Instance() *JWTService {
	if jwtService != nil {
		return jwtService
	}

	jwtService = &JWTService{}

	return jwtService
}

func (j *JWTService) GenerateToken(secretKey string, secretKeyRefreshToken string, value any) (string, string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	_secretKey := []byte(secretKey)
	_secretKeyRefreshToken := []byte(secretKeyRefreshToken)

	claims := &Claims{
		Value: value,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(_secretKey)
	if err != nil {
		return "", "", err
	}

	refreshExpirationTime := time.Now().Add(24 * time.Hour)
	refreshClaims := &Claims{
		Value: value,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(_secretKeyRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (j *JWTService) ValidateToken(secretKey string, tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
