package models

import "math"

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int   `json:"totalPages"`
}

func NewPagination(page, pageSize int, totalCount int64) Pagination {
	return Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: int(math.Ceil(float64(totalCount) / float64(pageSize))),
	}
}