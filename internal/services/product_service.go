package services

import (
	"boilerplate_go/domain"
	"boilerplate_go/internal/repository"
	"context"
)

// type IProductService interface {
// 	GetAllProducts(ctx context.Context) ([]domain.ProductVModel, error)
// }

// type ProductService struct {
// 	productService IProductService
// }

// func InitProductService(productService IProductService) *ProductService {
// 	return &ProductService{productService: productService}
// }

func GetAllProducts(ctx context.Context) ([]domain.ProductVModel, error) {
	products, _ := repository.ProductModel().GetAll(ctx)

	var payload []domain.ProductVModel
	for _, product := range products {
		payload = append(payload, *product.MapToVModel())
	}

	return payload, nil
}

func CreateProduct(ctx context.Context, vmodel *repository.ProductVModel) error {
	model := vmodel.MapToVModel()
	return model.Create(ctx)
}