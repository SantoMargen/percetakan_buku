package notification

type RequestNotification struct {
	KeyNotif    int    `json:"key_notif"`
	DescNotif   string `json:"desc_notif"`
	TitleNotif  string `json:"title_notif"`
	Receiver    int    `json:"receiver"`
	UrlRedirect string `json:"url_redirect"`
}

type RequestNotificationById struct {
	ID int `json:"id"`
}

type SentRequestNotification struct {
	RequestNotification
	Sender int `json:"sender"`
}

type NotificationPagination struct {
	Page   int                 `json:"page"`
	Size   int                 `json:"size"`
	Filter *FilterNotification `json:"filter"`
}

type FilterNotification struct {
	UserID int `json:"user_id"`
}
