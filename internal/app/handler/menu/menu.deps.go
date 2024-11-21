package menu

import (
	"context"
	"siap_app/internal/app/entity/menu"
)

type menuUC interface {
	GetMenu(ctx context.Context, role string) ([]menu.ResponseMenu, error)
}
