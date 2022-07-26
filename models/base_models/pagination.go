package base_models

import (
	"math"

	"github.com/bagastri07/be-test-kumparan/constants"
)

type PaginationParams struct {
	Page  uint64 `query:"page"`
	Limit uint64 `query:"limit"`
}

func (p *PaginationParams) ProcessLimitation() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit > constants.MaxLimit {
		p.Limit = constants.MaxLimit
	}
	if p.Limit <= constants.MinLimit {
		p.Limit = constants.MinLimit
	}
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
