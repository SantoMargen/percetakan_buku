package category

type RequestCategory struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	EntryUser    string `json:"entry_user"`
}

type RequestCategoryByID struct {
	ID int `json:"id"`
}

type RequestCategoryUpdate struct {
	ID int `json:"id"`
	RequestCategory
}

type PaginationCategory struct {
	Page   int             `json:"page"`
	Size   int             `json:"size"`
	Filter *FilterCategory `json:"filter"`
}

type FilterCategory struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
