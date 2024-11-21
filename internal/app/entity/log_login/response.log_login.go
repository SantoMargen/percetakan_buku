package loglogin

import "time"

type LogloginResponse struct {
	ID          int        `db:"id" json:"id"`
	Email       string     `db:"email" json:"email"`
	FullName    string     `db:"full_name" json:"full_name"`
	Role        string     `db:"role" json:"role"`
	IPAddress   string     `db:"ip_address" json:"ip_address"`
	LoginTime   time.Time  `db:"login_time" json:"login_time"`
	LogoutTime  *time.Time `db:"logout_time" json:"logout_time,omitempty"`
	ProcessTime *time.Time `db:"process_time" json:"process_time,omitempty"`
	ExpiredTime *time.Time `db:"expired_time" json:"expired_time,omitempty"`
}
