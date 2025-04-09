package usecase

import (
	"boilerplate_go/internal/dto"
	automapper "boilerplate_go/internal/helper"
	"boilerplate_go/internal/model"
	"boilerplate_go/internal/repository"
	"boilerplate_go/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type ProductUseCase interface {
	Search(ctx context.Context, searchModel dto.SearchProductDto) (dto.BaseResponse, error)
	GetByID(id string) (*dto.ProductDto, error)
	GetAll() ([]dto.ProductDto, error)
	Create(productDto *dto.ProductDto) error
	Update(productDto *dto.ProductDto, id string) error
	Delete(id string) error
}

type productUseCase struct {
	repo repository.BaseRepository[model.Product]
}

func NewProductUseCase(repo repository.BaseRepository[model.Product]) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (productUseCase *productUseCase) Search(ctx context.Context, searchModel dto.SearchProductDto) (dto.BaseResponse, error) {
	// retrieve products from cache if available
	cacheKey := fmt.Sprintf("products:search:keyword=%s:pageNumber=%d:pageSize=%d",
        searchModel.Keyword,
        searchModel.PageNumber,
        searchModel.PageSize,
    )

	cached, _ := utils.RedisClient.Get(ctx, cacheKey)
	if cached != "" {
        var payload dto.BaseResponse
        if err := json.Unmarshal([]byte(cached), &payload); err == nil {
            return payload, nil
        }
    }

	keyword := "%" + searchModel.Keyword + "%"
	entities, total, err := productUseCase.repo.Search("name LIKE ? OR description LIKE ?", searchModel.PageNumber, searchModel.PageSize, keyword, keyword)

	if err != nil {
		return dto.NewBaseResponse(500, err.Error(), dto.ProductsResponse{
			Records: nil,
			Pagination: dto.NewPagination(searchModel.PageNumber, searchModel.PageSize, 0),
		}), err
	}

	dtoProducts := make([]dto.ProductDto, len(entities))
	for i, entity := range entities {
		dto := dto.ProductDto{}
		automapper.MapLoose(entity, &dto)
		dtoProducts[i] = dto
	}

	result := dto.NewBaseResponse(200, "success", dto.ProductsResponse{
		Records: dtoProducts,
		Pagination: dto.NewPagination(searchModel.PageNumber, searchModel.PageSize, total),
	})

	// Set payload to Redis
	data, _ := json.Marshal(result)
	if err:= utils.RedisClient.Set(ctx, cacheKey, string(data)); err != nil {
		return result, err
	}

	return result, nil
}

// GetAll implements ProductUseCase.
func (p *productUseCase) GetAll() ([]dto.ProductDto, error) {
	entities, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	dtoProducts := make([]dto.ProductDto, len(entities))
	for i, entity := range entities {
		dto := dto.ProductDto{}
		automapper.MapLoose(entity, &dto)
		dtoProducts[i] = dto
	}

	return dtoProducts, nil
}

// GetByID implements ProductUseCase.
func (p *productUseCase) GetByID(id string) (*dto.ProductDto, error) {
	dto := dto.ProductDto{}
	product, err := p.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	automapper.MapLoose(product, &dto)
	return &dto, nil
}

// Create implements ProductUseCase.
func (p *productUseCase) Create(productDto *dto.ProductDto) error {
	var entity model.Product
	automapper.MapLoose(productDto, &entity)
	return p.repo.Create(&entity)
}

// Update implements ProductUseCase.
func (p *productUseCase) Update(productDto *dto.ProductDto, id string) error {
	entity, _ := p.repo.GetByID(id)
	if entity == nil {
		return errors.New("product not found")
	}

	newEntity := model.Product{}
	automapper.MapLoose(productDto, &newEntity)

	return p.repo.Update(&newEntity, id)
}

// Delete implements ProductUseCase.
func (p *productUseCase) Delete(id string) error {
	entity, _ := p.repo.GetByID(id)
	if entity == nil {
		return errors.New("product not found")
	}

	return p.repo.Delete(id)
}
