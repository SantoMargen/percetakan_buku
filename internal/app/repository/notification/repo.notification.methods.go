package notification

import (
	"context"
	"database/sql"
	"fmt"
	"siap_app/internal/app/entity/notification"
	"strconv"

	"github.com/pkg/errors"
)

func (r *repository) GetNotificationAll(ctx context.Context, input notification.NotificationPagination, role string) ([]notification.ResponseNotification, int64, error) {

	var (
		noficationList []notification.ResponseNotification
		offset         int
		query          string
		countQuery     string
		total          int64
	)

	offset = (input.Page - 1) * input.Size

	query = "SELECT " + columnSelectNotification + " FROM log_notif WHERE 1=1"
	countQuery = "SELECT COUNT(*) FROM log_notif WHERE 1=1"

	var args []interface{}
	var nextLimit int

	if role == "ADMIN" {
		if input.Filter != nil {
			if input.Filter.UserID != 0 {
				query += " AND receiver = $" + strconv.Itoa(len(args)+1)
				countQuery += " AND receiver = $" + strconv.Itoa(len(args)+1)
				args = append(args, input.Filter.UserID)
				nextLimit++
			}
		}
	} else {
		query += " AND receiver = $" + strconv.Itoa(len(args)+1)
		countQuery += " AND receiver = $" + strconv.Itoa(len(args)+1)
		args = append(args, input.Filter.UserID)
		nextLimit++
	}

	query += " ORDER BY id_log_notif ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataNotification notification.ResponseNotification
		if err := rows.Scan(
			&dataNotification.IDNotif,
			&dataNotification.KeyNotification,
			&dataNotification.DescNotif,
			&dataNotification.TitleNotif,
			&dataNotification.Sender,
			&dataNotification.Receiver,
			&dataNotification.FlagRead,
			&dataNotification.UrlRedirect,
			&dataNotification.CreatedTime,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		noficationList = append(noficationList, dataNotification)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return noficationList, total, nil
}

func (r *repository) CreateLogNotif(ctx context.Context, input notification.SentRequestNotification) error {
	_, err := r.db.ExecContext(ctx, qryInsertNotification,
		input.KeyNotif,
		input.DescNotif,
		input.TitleNotif,
		input.Sender,
		input.Receiver,
		input.UrlRedirect,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create notification")
	}

	return nil
}

func (r *repository) DeleteNotification(ctx context.Context, id int, userId int) error {
	_, err := r.db.ExecContext(ctx, queryDeleteNotification,
		id,
		userId,
	)

	if err != nil {
		return errors.Wrap(err, "failed to delete notification")
	}

	return nil
}

func (r *repository) GetNotificationById(ctx context.Context, notifId int, userdID int) (notification.ResponseNotification, error) {
	var dataNotification notification.ResponseNotification
	err := r.db.QueryRowContext(ctx, queryNotificationById, notifId, userdID).Scan(
		&dataNotification.IDNotif,
		&dataNotification.KeyNotification,
		&dataNotification.DescNotif,
		&dataNotification.TitleNotif,
		&dataNotification.Sender,
		&dataNotification.Receiver,
		&dataNotification.FlagRead,
		&dataNotification.UrlRedirect,
		&dataNotification.CreatedTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return dataNotification, errors.Wrap(err, "notification not found")
		}
		return dataNotification, errors.Wrap(err, "failed to get notification by id")
	}

	return dataNotification, nil
}
