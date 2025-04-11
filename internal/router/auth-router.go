package router

import (
	"boilerplate_go/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Group, a controller.AuthController) {
	authGroup := e.Group("/auth")

	authGroup.POST("/register", a.Register)
	authGroup.POST("/login", a.Login)
	// authGroup.POST("/logout", a.Logout)
	authGroup.GET("/me", a.Me)
	// authGroup.POST("/refresh-token", a.RefreshToken)
}
