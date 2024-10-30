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
