package common

type PageRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	SortBy   string `json:"sort_by,omitempty"`
	Order    string `json:"order,omitempty"`
}

type PageResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}
