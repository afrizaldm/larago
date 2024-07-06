package jwt

import (
	"errors"
	"fmt"
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

func (j *JWTService) GenerateToken(value any, secretKey string, expirationTimeMinute int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationTimeMinute) * time.Minute)

	_secretKey := []byte(secretKey)

	claims := &Claims{
		Value: value,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(_secretKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (j *JWTService) GenerateTokens(value any, secretKeys ...string) ([]string, error) {
	if len(secretKeys) == 0 {
		return nil, fmt.Errorf("no secret keys provided")
	}

	var tokens []string
	for _, secretKey := range secretKeys {
		token, err := jwtService.GenerateToken(value, secretKey, 1440)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
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
