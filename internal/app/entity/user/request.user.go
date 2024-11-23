package user

type RegisterRequest struct {
	FullName    string  `json:"full_name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	PhoneNumber *string `json:"phone_number"`
	Role        *string `json:"role"`
	Gender      string  `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DataToken struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type RegisterByAdminRequest struct {
	FullName    string  `json:"full_name"`
	Email       string  `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber string  `json:"phone_number"`
	Role        string  `json:"role"`
	Gender      string  `json:"gender"`
}

type UpdateRoleRequest struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type UpdatePaswordeRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
