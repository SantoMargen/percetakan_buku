package entity

type Response[T any] struct {
	Data T     `json:"data"`
	Page *Page `json:"page,omitempty"`
}

type Page struct {
	Size       int    `json:"size"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Current    int    `json:"current"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type ErrResponse struct {
	ResponseCode int    `json:"responseCode"`
	ErrType      string `json:"err_type"`
	ErrMessage   string `json:"err_message"`
	Status       string `json:"status"`
}

type contextKey string

const (
	UserIDKey       contextKey = "userId"
	FullNameKey     contextKey = "fullName"
	RoleKey         contextKey = "role"
	EmailKey        contextKey = "email"
	IsAuthorizedKey contextKey = "isAuthorized"
)

type RequestData struct {
	Data string `json:"data"`
}

type RequestInput struct {
	Request interface{} `json:"request"`
}
