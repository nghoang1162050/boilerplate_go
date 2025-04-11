package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CasbinMiddleware(e *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user").(*jwt.Token)
			claims := userToken.Claims.(jwt.MapClaims)
			username := claims["username"].(string)

			obj := c.Path()
			act := c.Request().Method

			ok, err := e.Enforce(username, obj, act)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Casbin error")
			}
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
			}
			return next(c)
		}
	}
}

