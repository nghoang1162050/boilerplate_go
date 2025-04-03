package models

import (
	"context"
)

type Product struct {
	Id    int     `gorm:"column:id;type:"int";primaryKey"`
	Name  string  `gorm:"column:name;type:nvarchar(45)"`
	Price float64 `gorm:"column:price;type:decimal(10,2)"`
}

var product *Product = &Product{}

func ProductModel() *Product {
	return product
}

func (model *Product) MapToVModel() *ProductVModel {
	return &ProductVModel{
		Id: model.Id,
		Name: model.Name,
		Price: model.Price,
	}
}

func (model *Product) GetAll(ctx context.Context) (models []Product, err error) {
	result := db.Debug().Model(model).WithContext(ctx).Order("id DESC").Find(&models)
	return models, result.Error
}

func (model *Product) Create(ctx context.Context) error {
	return db.Model(model).WithContext(ctx).Create(&model).Error
}