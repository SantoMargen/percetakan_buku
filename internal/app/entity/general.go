package entity

type Response[T any] struct {
	Data T `json:"data"`
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

type ResponsePagination struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type TokenData struct {
	UserId       int
	FullName     string
	Role         string
	Email        string
	IsAuthorized bool
}
