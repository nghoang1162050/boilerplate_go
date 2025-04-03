package products

import (
	"boilerplate_go/internal/db/models"
	"boilerplate_go/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	product := models.ProductVModel{}
	if err := c.Bind(&product); err != nil {
		return err
	}

	model := product.MapToModel()
	if err := model.Create(c.Request().Context()); err != nil {
		return c.JSON(500, handlers.Error("500", err))
	}

	return c.JSON(201, handlers.Accepted())
}