package products

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/handlers"
	constants "boilerplate_go/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, handlers.Error("invalid ID parameter", err))
	}

	product, err := models.ProductModel().GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}
	if product == nil {
		return c.JSON(404, handlers.Error(constants.MSG_RECORD_NOT_FOUND, nil))
	}

	if err := product.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}

	return c.NoContent(204)
}