package middlewares

import (
	constants "boilerplate_go/pkg/utils"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip authentication for login
			if c.Path() == "/login" || c.Path() == "/swagger/*" {
				return next(c)
			}

			headerVal := c.Request().Header.Get("Authorization")
			if headerVal == "" {
				return echo.ErrUnauthorized
			}
			
			fields := strings.Fields(headerVal)
			if len(fields) < 2 || strings.ToLower(fields[0]) != "bearer" {
                return echo.ErrUnauthorized
            }
            tokenString := fields[1]

			claims := &jwt.RegisteredClaims{}
			// TODO: verify the token with the public key and particularly the issuer
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
                return []byte(constants.JWT_SECRET_KEY), nil
            })
            if err != nil || !token.Valid {
                return echo.ErrUnauthorized
            }

			return next(c)
		}
	}
}