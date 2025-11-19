package dto

type SearchResponse struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Results  interface{} `json:"results"`
}
