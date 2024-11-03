package user

import "time"

type ResponseUser struct {
	ID           int        `json:"id"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Role         string     `json:"role"`
	Gender       string     `json:"gender"`
	PhoneNumber  *string    `json:"phone_number"`
	TanngalLahir *time.Time `json:"tanggal_lahir"`
	ImageProfile *string    `json:"imageProfile"`
	Address      *string    `json:"address"`
	City         *string    `json:"city"`
	Country      *string    `json:"country"`
	CreatedBy    *int       `json:"created_by"`
	UpdatedBy    *int       `json:"updated_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
