package products

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/handlers"
	"boilerplate_go/pkg/api/helpers"
	constants "boilerplate_go/pkg/utils"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	product := models.ProductVModel{}

	if err := c.Bind(&product); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(product); err != nil {
		return c.JSON(400, handlers.ValidationErrors(err))
	}
	
	model := product.MapToModel()
	if err := model.Create(c.Request().Context()); err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}

	return c.JSON(201, handlers.Accepted())
}