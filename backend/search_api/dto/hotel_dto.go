package dto

import "search_api/domain"

type SearchResponse struct {
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Results  []domain.Hotel `json:"results"`
}
