package levelusers

import (
	"context"
	levelusers "siap_app/internal/app/entity/level_users"
)

type levelUserRepo interface {
	GetLevelUsers(ctx context.Context) ([]levelusers.LevelUser, error)
	GetLevelUserByID(ctx context.Context, id int) (levelusers.LevelUser, error)
}
