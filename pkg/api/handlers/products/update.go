package products

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/handlers"
	"boilerplate_go/pkg/api/helpers"
	constants "boilerplate_go/pkg/utils"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
        return c.JSON(400, handlers.Error("invalid ID parameter", err))
    }

	current, err := models.ProductModel().GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}
	if current == nil {
		return c.JSON(404, handlers.Error(constants.MSG_RECORD_NOT_FOUND, nil))
	}

	product := models.ProductUpdateVModel{}

	if err := c.Bind(&product); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(product); err != nil {
		return c.JSON(400, handlers.ValidationErrors(err))
	}
	
	model := product.MapToModel(id)
	fmt.Printf("model: %v", model)

	if err := model.Update(c.Request().Context()); err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}

	return c.JSON(200, handlers.Success(model.MapToVModel()))
}