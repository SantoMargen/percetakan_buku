package levelusers

import (
	"context"
	levelusers "siap_app/internal/app/entity/level_users"
)

type levelUserUC interface {
	GetLevelUsers(ctx context.Context) ([]levelusers.LevelUser, error)
	GetLevelUsersByID(ctx context.Context, id int) (*levelusers.LevelUser, error)
}
