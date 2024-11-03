package category

type ResponseCategory struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	EntryUser    string `json:"entry_user"`
	EntryTime    string `json:"entry_time"`
}
