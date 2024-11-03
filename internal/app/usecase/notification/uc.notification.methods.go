package notification

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/notification"
)

func (uc *UseCase) GetNotificationAll(ctx context.Context, input notification.NotificationPagination, role string) ([]notification.ResponseNotification, int64, error) {

	resp, total, err := uc.notificationRepo.GetNotificationAll(ctx, input, role)

	if err != nil {
		return nil, 0, fmt.Errorf("error get data notification : %w", err)
	}

	return resp, total, nil

}

func (uc *UseCase) CreateLogNotif(ctx context.Context, input notification.SentRequestNotification) error {

	return uc.notificationRepo.CreateLogNotif(ctx, input)
}

func (uc *UseCase) DeleteNotification(ctx context.Context, id int, userID int) error {
	getCategoryResponse, err := uc.notificationRepo.GetNotificationById(ctx, id, userID)

	if err != nil {
		return err
	}

	if getCategoryResponse.IDNotif == "" {
		return fmt.Errorf("fetch notification failed")
	}

	return uc.notificationRepo.DeleteNotification(ctx, id, userID)

}
