package notification

type ResponseNotification struct {
	IDNotif         string `json:"id_notif"`
	KeyNotification string `json:"key_notif"`
	UrlRedirect     string `json:"url_redirect"`
	DescNotif       string `json:"desc_notif"`
	TitleNotif      string `json:"title_notif"`
	Sender          string `json:"sender_id"`
	Receiver        string `json:"receiver_id"`
	CreatedTime     string `json:"created_time"`
	FlagRead        string `json:"flag_read"`
}
