package loglogin

import "time"

type LogloginRequest struct {
	Email       string     `json:"email"`
	FullName    string     `json:"full_name"`
	Role        string     `json:"role"`
	IPAddress   string     `json:"ip_address"`
	LoginTime   time.Time  `json:"login_time"`
	LogoutTime  *time.Time `json:"logout_time,omitempty"`
	ProcessTime *time.Time `json:"process_time,omitempty"`
	ExpiredTime *time.Time `json:"expired_time,omitempty"`
}
