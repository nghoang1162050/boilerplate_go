package router

import (
	"boilerplate_go/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewProductRouter(e *echo.Group, p controller.ProductController) {
	productsGroup := e.Group("/products")

	productsGroup.GET("", p.Search)
    productsGroup.GET("/:id", p.GetByID)
    productsGroup.POST("", p.Create)
    productsGroup.PUT("/:id", p.Update)
    productsGroup.DELETE("/:id", p.Delete)
}
