package base_models

import (
	"math"
)

type PaginationParams struct {
	Page  uint64 `query:"page"`
	Limit uint64 `query:"limit"`
}

type PaginationResponse struct {
	Page       uint64      `json:"page"`
	Limit      uint64      `json:"limit"`
	TotalItems uint64      `json:"totalItems"`
	TotalPages uint64      `json:"totalPages"`
	Items      interface{} `json:"items"`
}

func (p *PaginationResponse) ToResponse(params *PaginationParams, items interface{}, totalItems uint64) {
	p.Page = params.Page
	p.Limit = params.Limit
	p.TotalItems = totalItems
	p.Items = items

	if totalItems != 0 || params.Limit != 0 {
		p.TotalPages = uint64(math.Ceil(float64(p.TotalItems) / float64(p.Limit)))
	} else {
		p.TotalPages = 0
	}

}
