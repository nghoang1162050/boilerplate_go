package utils

import (
	"os"
	"strconv"
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
