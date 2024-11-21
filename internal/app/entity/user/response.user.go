package user

import "time"

type ResponseUser struct {
	ID            int        `json:"id"`
	FullName      string     `json:"full_name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	Role          string     `json:"role"`
	Gender        string     `json:"gender"`
	PhoneNumber   *string    `json:"phone_number"`
	TanngalLahir  *time.Time `json:"tanggal_lahir"`
	ImageProfile  *string    `json:"imageProfile"`
	Address       *string    `json:"address"`
	City          *string    `json:"city"`
	Country       *string    `json:"country"`
	CreatedBy     *int       `json:"created_by"`
	UpdatedBy     *int       `json:"updated_by"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Status        int        `json:"status"`
	BornPlace     string     `json:"born_place"`
	Biography     string     `json:"biography"`
	Experience    string     `json:"experience"`
	Achievement   string     `json:"achievement"`
	LastEduaction string     `json:"last_education"`
	ResponseCard
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type ResponseCard struct {
	TotalUsers       int `json:"total_user"`
	ActiveUsers      int `json:"active_user"`
	NewUsersLastWeek int `json:"new_user_last_week"`
}

type ResponseLog struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	IPAddress  string `json:"ip_address"`
	UserAgent  string `json:"user_agent"`
	LoginTime  string `json:"login_time"`
	LogoutTime string `json:"logout_time"`
}
