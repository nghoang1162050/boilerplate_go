package products

import (
	"boilerplate_go/internal/db/models"
	"boilerplate_go/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	products, _ := models.ProductModel().GetAll(c.Request().Context())

	var payload []models.ProductVModel
	for _, product := range products {
		payload = append(payload, *product.MapToVModel())
	}

	return c.JSON(200, handlers.Success(payload))
}
