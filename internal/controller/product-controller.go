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

// Search implements ProductController.
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

func NewProductController(productUseCase usecase.ProductUseCase) ProductController {
	return &productController{productUseCase: productUseCase}
}

// GetAll implements ProductController.
func (p *productController) GetAll(ctx echo.Context) error {
	products, err := p.productUseCase.GetAll()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}

	return ctx.JSON(200, products)
}

// GetByID implements ProductController.
func (p *productController) GetByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	product, err := p.productUseCase.GetByID(idParam)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, product)
}

// Create implements ProductController.
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
func (p *productController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := p.productUseCase.Delete(id); err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.NoContent(204)
}
