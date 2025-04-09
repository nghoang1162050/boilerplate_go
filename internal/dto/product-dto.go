package dto

import "time"

type ProductDto struct {
    ID            string    `json:"id" validate:"required,uuid4"`
    Name          string    `json:"name" validate:"required"`
    Brand         string    `json:"brand" validate:"required"`
    Description   string    `json:"description" validate:"required"`
    Price         float64   `json:"price" validate:"required,gt=0"`
    StockQuantity int32     `json:"stock_quantity" validate:"required,gte=0"`
    CPU           string    `json:"cpu" validate:"omitempty"`
    RAM           string    `json:"ram" validate:"omitempty"`
    Storage       string    `json:"storage" validate:"omitempty"`
    Gpu           string    `json:"gpu" validate:"omitempty"`
    CreatedAt     time.Time `json:"created_at" validate:"required"`
}

type SearchProductDto struct {
    Keyword string `query:"keyword"`
    PageNumber    int    `query:"pageNumber" validate:"required,gte=1"`
    PageSize   int    `query:"pageSize" validate:"required,gte=1,lte=100"`
}

type ProductsResponse struct {
    Records []ProductDto `json:"records"`
    Pagination Pagination `json:"pagination"`
}
