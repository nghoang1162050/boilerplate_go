package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
    // List of endpoints to ignore.
    ignoredPaths := map[string]bool{
        "/api/auth/register": true,
        "/api/auth/login":    true,
    }

    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // If the request path is in the ignored list, simply pass to the next handler.
            if ignoredPaths[c.Path()] {
                return next(c)
            }

            authHeader := c.Request().Header.Get("Authorization")
            if authHeader == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
            }

            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
                return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header")
            }

            tokenString := parts[1]
            secret := os.Getenv("JWT_SECRET")
            claims := &jwt.RegisteredClaims{}
            token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })
            
            if err != nil || !token.Valid {
                return echo.ErrUnauthorized
            }

            return next(c)
        }
    }
}
