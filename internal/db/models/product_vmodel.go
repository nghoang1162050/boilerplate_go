package models

type ProductVModel struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (model *ProductVModel) MapToModel() *Product {
	return &Product{
		Id:    model.Id,
		Name:  model.Name,
		Price: model.Price,
	}
}