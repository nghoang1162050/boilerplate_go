package products

import (
	"boilerplate_go/internal/cache"
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/handlers"
	"boilerplate_go/pkg/api/helpers"
	constants "boilerplate_go/pkg/utils"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAll godoc
// @Summary Get all products
// @Description Retrieves a list of products with pagination
// @Tags products
// @Accept json
// @Produce json
// @Param name query string false "Search by name"
// @Param price query number false "Filter by price"
// @Param pageNumber query int true "Page number"
// @Param pageSize query int true "Page size"
// @Success 200 {object} models.ProductsVModel
// @Router /products [get]
func GetAll(c echo.Context) error {
	searchModel := models.SearchProductVModel{}

	if err := c.Bind(&searchModel); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(searchModel); err != nil {
		return c.JSON(400, handlers.ValidationErrors(err))
	}

	ctx := c.Request().Context()
    cacheKey := fmt.Sprintf("products:list:name=%s:price=%v:page=%d:pageSize=%d",
        searchModel.Name,
        searchModel.Price,
        searchModel.PageNumber,
        searchModel.PageSize,
    )
    cached, _ := cache.RedisClient.Get(ctx, cacheKey)
    if cached != "" {
        var payload models.ProductsVModel
        if err := json.Unmarshal([]byte(cached), &payload); err == nil {
            return c.JSON(200, handlers.Success(payload))
        }
    }

	products, total, err := models.ProductModel().GetAll(c.Request().Context(), &searchModel)
	if err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}		

	payload := models.ProductsVModel{
		Products: make([]models.ProductVModel, len(products)),
		Pagination: models.NewPagination( 
			searchModel.PageNumber,
			searchModel.PageSize,
			total,
		),
	}

	for i, p := range products {
		payload.Products[i] = *p.MapToVModel()
	}

	// Set payload to Redis
	data, _ := json.Marshal(payload)
	if err:= cache.RedisClient.Set(ctx, cacheKey, string(data)); err != nil {
		return c.JSON(500, handlers.Error(constants.MSG_INTERNAL_SERVER, err))
	}

	return c.JSON(200, handlers.Success(payload))
}

func GetById(c echo.Context) error {
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

	return c.JSON(200, handlers.Success(product.MapToVModel()))
}
