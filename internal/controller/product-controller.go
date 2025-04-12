package controller

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ProductController interface {
	Search(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) ProductController {
	return &productController{productUseCase: productUseCase}
}

// Search implements ProductController.
// @Summary Search Products
// @Description Retrieves products based on search query, page number and page size.
// @Tags Products
// @Accept json
// @Produce json
// @Param keyword query string false "Search Keyword"
// @Param pageNumber query int true "Page Number (>=1)"
// @Param pageSize query int true "Page Size (>=1)"
// @Success 200 {object} dto.BaseResponse "Success"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products [get]
func (p *productController) Search(ctx echo.Context) error {
	var searchModel dto.SearchProductDto
	if err := ctx.Bind(&searchModel); err != nil {
		return ctx.JSON(400, err.Error())
	}

	result, err := p.productUseCase.Search(ctx.Request().Context(), searchModel)
	if err != nil {
		return ctx.JSON(500, result)
	}

	return ctx.JSON(200, result)
}

// GetAll implements ProductController.
// @Summary List All Products
// @Description Retrieves a list of all products.
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} dto.ProductDto "List of products"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products/all [get]
func (p *productController) GetAll(ctx echo.Context) error {
	products, err := p.productUseCase.GetAll()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}

	return ctx.JSON(200, products)
}

// GetByID implements ProductController.
// @Summary Get Product By ID
// @Description Retrieves a product by its identifier.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ProductDto "Product found"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products/{id} [get]
func (p *productController) GetByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	product, err := p.productUseCase.GetByID(idParam)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, product)
}

// Create implements ProductController.
// @Summary Create Product
// @Description Creates a new product with the provided payload.
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.ProductDto true "Product Data"
// @Success 201 {object} dto.ProductDto "Product created"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products [post]
func (p *productController) Create(ctx echo.Context) error {
	var productDto dto.ProductDto

	if err := ctx.Bind(&productDto); err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := p.productUseCase.Create(&productDto); err != nil {
		return ctx.JSON(500, err.Error())
	}

	return ctx.JSON(201, productDto)
}

// Update implements ProductController.
// @Summary Update Product
// @Description Updates an existing product with the specified ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body dto.ProductDto true "Updated Product Data"
// @Success 200 {object} dto.ProductDto "Product updated"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products/{id} [put]
func (p *productController) Update(ctx echo.Context) error {
	var productDto dto.ProductDto

	if err := ctx.Bind(&productDto); err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := p.productUseCase.Update(&productDto, ctx.Param("id")); err != nil {
		return ctx.JSON(500, err.Error())
	}

	return ctx.JSON(200, productDto)
}

// Delete implements ProductController.
// @Summary Delete Product
// @Description Deletes a product with the specified ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /products/{id} [delete]
func (p *productController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := p.productUseCase.Delete(id); err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.NoContent(204)
}
