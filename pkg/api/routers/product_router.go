package routers

import (
	"boilerplate_go/pkg/api/handlers/products"

	"github.com/labstack/echo/v4"
)

func InitProductRouter(e *echo.Echo) {
	e.GET("/products", products.GetAll)
	e.GET("/products/:id", products.GetById)
	e.POST("/products", products.Create)
	e.PUT("/products/:id", products.Update)
	e.DELETE("/products/:id", products.Delete)
}