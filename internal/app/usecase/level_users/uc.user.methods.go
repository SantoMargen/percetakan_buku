package levelusers

import (
	"context"
	levelusers "siap_app/internal/app/entity/level_users"
)

func (uc *UseCase) GetLevelUsers(ctx context.Context) ([]levelusers.LevelUser, error) {
	var (
		data = []levelusers.LevelUser{}
		err  error
	)
	data, err = uc.levelUserRepo.GetLevelUsers(ctx)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (uc *UseCase) GetLevelUsersByID(ctx context.Context, id int) (*levelusers.LevelUser, error) {
	var (
		data = levelusers.LevelUser{}
		err  error
	)

	data, err = uc.levelUserRepo.GetLevelUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
