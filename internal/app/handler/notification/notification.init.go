package notification

type Handler struct {
	notificationUC notificationUC
}

func New(notificationUC notificationUC) *Handler {
	return &Handler{notificationUC: notificationUC}
}
