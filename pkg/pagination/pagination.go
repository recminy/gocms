package pagination

import "gorm.io/gorm"

type Pagination struct {
	CurrentPage int `json:"current_page"`
	From int `json:"from"`
	To int `json:"to"`
	LastPage int `json:"last_page"`
	PrevPageUrl string `json:"prev_page_url"`
	NextPageUrl string `json:"next_page_url"`
	Total int `json:"total"`
	PerPage int `json:"per_page"`
	db *gorm.DB
}

func (p Pagination) Paging() {

}