package status

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/status"
)

func (uc *UseCase) GetStatusAll(ctx context.Context, input status.PaginationStatus) ([]status.StatusResponse, int64, error) {

	resp, total, err := uc.statusRepo.GetStatusAll(ctx, input)

	if err != nil {
		return nil, 0, fmt.Errorf("error get data publishers : %w", err)
	}

	return resp, total, nil

}
