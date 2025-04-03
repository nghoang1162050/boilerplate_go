package rest

import (
	"boilerplate_go/internal/repository"
	"boilerplate_go/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)


func NewProductHandler(e *echo.Echo) {
	e.GET("/products", GetAllProducts)
	e.POST("/products", Post)
}

func GetAllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	products, _ := services.GetAllProducts(ctx)
	return c.JSON(200, products)
}

func Post(c echo.Context) error {
	vmodel := &repository.ProductVModel{}

	if err := c.Bind(vmodel); err != nil {
		return c.JSON(http.StatusBadRequest, vmodel)
	}
	ctx := c.Request().Context()

	return services.CreateProduct(ctx, vmodel)
}