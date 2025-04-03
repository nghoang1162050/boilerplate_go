package repository

type ProductVModel struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (model *ProductVModel) MapToVModel() *Product {
	return &Product{
		Id:    model.Id,
		Name:  model.Name,
		Price: model.Price,
	}
}