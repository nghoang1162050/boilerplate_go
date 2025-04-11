package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWTSecret(username string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	jwt_expired, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	expired := time.Now().Add(time.Duration(jwt_expired) * time.Minute).Unix()

	claims := &jwt.RegisteredClaims{
		Issuer:    username,
		Subject:  username,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Unix(expired, 0)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, expired, err
}

func ExtractUsernameFromToken(header string) (string, error) {
	tokenString, err := ExtractTokenFromHeader(header)
	secret := os.Getenv("JWT_SECRET")
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.Subject, nil
}

func ExtractTokenFromHeader(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid Authorization header")
	}
	return parts[1], nil
}
