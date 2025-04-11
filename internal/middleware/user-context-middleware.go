package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func UserContextMiddleware() echo.MiddlewareFunc {
	jwtMw := JWTMiddleware()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return jwtMw(func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(jwt.MapClaims)

			if username, ok := claims["username"].(string); ok {
				c.Set("username", username)
			}

			// if role, ok := claims["role"].(string); ok {
			// 	c.Set("role", role)
			// }

			return next(c)
		})
	}
}
