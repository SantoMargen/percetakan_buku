package menu

import (
	"context"
	"siap_app/internal/app/entity/menu"
)

func (r *repository) GetMenu(ctx context.Context) ([]menu.Menu, error) {
	var menuData []menu.Menu

	rows, err := r.db.QueryContext(ctx, queryGetMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m menu.Menu
		err := rows.Scan(
			&m.MenuID,
			&m.MenuParent,
			&m.Kategori,
			&m.MenuIcon,
			&m.MenuTitle,
			&m.MenuLink,
			&m.MenuActive,
			&m.LevelUser,
		)

		if err != nil {
			return nil, err
		}
		menuData = append(menuData, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return menuData, nil
}

func (r *repository) GetMenuKategory(ctx context.Context) ([]menu.Kategory, error) {
	var menuKategory []menu.Kategory

	rows, err := r.db.QueryContext(ctx, queryGetKategoryMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k menu.Kategory
		err := rows.Scan(
			&k.KategoriID,
			&k.Kategori,
			&k.MenuUrut,
		)

		if err != nil {
			return nil, err
		}
		menuKategory = append(menuKategory, k)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return menuKategory, nil
}
