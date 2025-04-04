package models

import (
	"context"
)

type Product struct {
	Id    int     `gorm:"column:id;type:int;primaryKey"`
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

func (model *Product) GetAll(ctx context.Context, searchModel *SearchProductVModel) (models []Product, total int64, err error) {
	dbQuery := db.Model(model).WithContext(ctx)

	if searchModel.Name != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+searchModel.Name+"%")
	}
	if searchModel.Price != 0 {
		dbQuery = dbQuery.Where("price = ?", searchModel.Price)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	dbQuery = dbQuery.Offset((searchModel.PageNumber - 1) * searchModel.PageSize).
        Limit(searchModel.PageSize).
        Order("id DESC").
        Select("id, name, price")

	if err := dbQuery.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	return models, total, nil
}

func (model *Product) Create(ctx context.Context) error {
	return db.Model(model).WithContext(ctx).Create(&model).Error
}

func (model *Product) GetById(ctx context.Context, id int) (*Product, error) {
	product := &Product{}
	if err := db.Model(model).WithContext(ctx).Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (model *Product) Update(ctx context.Context) error {
	return db.WithContext(ctx).Model(model).Updates(model).Error
}

func (model *Product) Delete(ctx context.Context, id int) error {
	return db.Model(model).WithContext(ctx).Where("id = ?", id).Delete(model).Error
}
