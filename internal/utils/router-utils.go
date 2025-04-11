package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func SyncRoutersToCasbin(e *echo.Echo, enforcer *casbin.Enforcer) {
	routes := e.Routes()

	for _, route := range routes {
		path := route.Path
		method := route.Method

		if ShouldIgnoreRequest(path) {
			continue // Skip ignored routes
		}

		// admin is default user for all routes
		hasPolicy, err := enforcer.HasPolicy("admin", path, method)
		if hasPolicy || err != nil {
			continue // Skip if the policy already exists
		}

		// Add the route to Casbin with a wildcard for the user ID
		enforcer.AddPolicy("admin", path, method)
	}
}