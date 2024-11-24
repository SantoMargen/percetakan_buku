package status

import (
	"context"
	"siap_app/internal/app/entity/status"
)

type statusUC interface {
	GetStatusAll(ctx context.Context, input status.PaginationStatus) ([]status.StatusResponse, int64, error)
}
