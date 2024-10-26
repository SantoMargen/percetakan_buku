package menu

import (
	"context"
	"siap_app/internal/app/entity/menu"
)

type menuRepo interface {
	GetMenuKategory(ctx context.Context) ([]menu.Kategory, error)
	GetMenu(ctx context.Context) ([]menu.Menu, error)
}
