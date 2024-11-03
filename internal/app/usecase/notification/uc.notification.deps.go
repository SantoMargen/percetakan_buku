package notification

import (
	"context"
	"siap_app/internal/app/entity/notification"
)

type notificationRepo interface {
	GetNotificationAll(ctx context.Context, input notification.NotificationPagination, role string) ([]notification.ResponseNotification, int64, error)
	CreateLogNotif(ctx context.Context, input notification.SentRequestNotification) error
	DeleteNotification(ctx context.Context, id int, userID int) error
	GetNotificationById(ctx context.Context, id int, userID int) (notification.ResponseNotification, error)
}
