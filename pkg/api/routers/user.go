package routers

import (
	"boilerplate_go/pkg/api/handlers/users"

	"github.com/labstack/echo/v4"
)

func InitUserRouters(e *echo.Echo) {
	e.POST("/login", users.Login)
}