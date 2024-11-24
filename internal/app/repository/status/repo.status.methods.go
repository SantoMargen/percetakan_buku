package status

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/status"
	"strconv"
)

func (r *repository) GetStatusAll(ctx context.Context, input status.PaginationStatus) ([]status.StatusResponse, int64, error) {

	var (
		dataStatusList []status.StatusResponse
		offset         int
		query          string
		countQuery     string
		total          int64
	)

	offset = (input.Page - 1) * input.Size

	query = "SELECT " + columnSelectStatus + " FROM status_submission WHERE 1=1"
	countQuery = "SELECT COUNT(*) FROM status_submission WHERE 1=1"

	var args []interface{}

	if input.Filter != nil {
		if input.Filter.IDStatus != "" {
			query += " AND id_status = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND id_status = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.IDStatus)
		}

	}

	query += " ORDER BY status_submission ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataStatus status.StatusResponse
		if err := rows.Scan(
			&dataStatus.IdStatus,
			&dataStatus.DescStatus,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataStatusList = append(dataStatusList, dataStatus)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataStatusList, total, nil
}
