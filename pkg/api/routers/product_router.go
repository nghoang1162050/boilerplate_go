package routers

import (
	"boilerplate_go/pkg/api/handlers/products"

	"github.com/labstack/echo/v4"
)

func InitProductRouter(e *echo.Echo) {
	e.GET("/products", products.GetAll)
	e.POST("/products", products.Create)
}