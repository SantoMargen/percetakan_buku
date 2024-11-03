package notification

import (
	"context"
	"siap_app/internal/app/entity/notification"
)

type notificationUC interface {
	GetNotificationAll(ctx context.Context, input notification.NotificationPagination, role string) ([]notification.ResponseNotification, int64, error)
	CreateLogNotif(ctx context.Context, input notification.SentRequestNotification) error
	DeleteNotification(ctx context.Context, id int, userId int) error
}
