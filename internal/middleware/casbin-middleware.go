package middleware

import (
	"boilerplate_go/internal/utils"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func CasbinMiddleware(e *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Ignore certain requests based on the path and method
			if utils.ShouldIgnoreRequest(c.Path()) {
                return next(c)
            }

			roles, _ := utils.ExtractRolesFromToken(c.Request().Header.Get("Authorization"))
			for _, role := range roles {
				path := c.Path()
				method := c.Request().Method
				if ok, _ := e.Enforce(role, path, method); ok {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
		}
	}
}
