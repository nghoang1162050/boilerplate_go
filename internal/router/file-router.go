package router

import (
	"boilerplate_go/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewFileRouter(e *echo.Group, f controller.FileController) {
	productsGroup := e.Group("/files")

	productsGroup.GET("", f.Search)
	productsGroup.GET("/download", f.Download)
	productsGroup.POST("", f.Upload)
	productsGroup.DELETE("", f.Delete)
}
