package models

type ProductVModel struct {
	Id    int     `json:"id" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price"`
}

func (model *ProductVModel) MapToModel() *Product {
	return &Product{
		Id:    model.Id,
		Name:  model.Name,
		Price: model.Price,
	}
}

type SearchProductVModel struct {
	Name       string  `query:"name"`
	Price      float64 `query:"price"`
	PageNumber int     `query:"pageNumber" validate:"required,min=1"`
	PageSize   int     `query:"pageSize" validate:"required,min=1,max=100"`
}

type ProductsVModel struct {
	Products   []ProductVModel `json:"products"`
	Pagination `json:"pagination"`
}

type ProductUpdateVModel struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price"`
}

func (model *ProductUpdateVModel) MapToModel(id int) *Product {
	return &Product{
		Id:    id,
		Name:  model.Name,
		Price: model.Price,
	}
}