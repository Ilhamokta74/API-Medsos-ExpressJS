package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key") // Ganti dengan secret key Anda

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken: Membuat JWT token
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam
	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken: Verifikasi JWT token
func VerifyToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
