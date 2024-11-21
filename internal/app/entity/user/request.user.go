package user

import "time"

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
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Gender      string `json:"gender"`
}

type UpdateRoleRequest struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type UpdatePaswordeRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type PaginationUser struct {
	Page   int         `json:"page"`
	Size   int         `json:"size"`
	Filter *FilterUser `json:"filter"`
}

type FilterUser struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type RequestByEmail struct {
	Email string `json:"email"`
}

type RequestById struct {
	UserId string `json:"userId"`
}

type RequestUpdateUser struct {
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	NewEmail     string `json:"email_new"`
	DateOfBirth  string `json:"date_of_birth"`
	BornPlace    string `json:"born_place"`
	Address      string `json:"address"`
	Graduated    string `json:"last_education"`
	Gender       string `json:"gender"`
	Biografi     string `json:"biography"`
	Experience   string `json:"experience"`
	Achievement  string `json:"achievement"`
	PhoneNumber  string `json:"phone_number"`
	Country      string `json:"country"`
	City         string `json:"city"`
	SetBirthDate time.Time
}

type PaginationLog struct {
	Page   int        `json:"page"`
	Size   int        `json:"size"`
	Filter *FilterLog `json:"filter"`
}

type FilterLog struct {
	Email        string `json:"email"`
	LastActivity string `json:"last_acitivity"`
}
