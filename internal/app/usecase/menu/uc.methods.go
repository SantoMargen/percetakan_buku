package menu

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/menu"
	"strings"
)

func (uc *UseCase) GetMenu(ctx context.Context, role string) ([]menu.ResponseMenu, error) {
	var (
		responseMenu = []menu.ResponseMenu{}
		menuFilter   []menu.MenuItem
		listKategori []menu.KategoryData
		err          error
	)

	menuKategory, err := uc.menuRepo.GetMenuKategory(ctx)
	if err != nil {
		return responseMenu, fmt.Errorf("failled to get data menu kategory")
	}

	menuData, err := uc.menuRepo.GetMenu(ctx)
	if err != nil {
		return responseMenu, fmt.Errorf("failled to get data menu")
	}

	for _, kategori := range menuKategory {
		tempKategory := menu.KategoryData{
			KategoriID: int(kategori.KategoriID),
			Kategori:   kategori.Kategori,
			MenuUrut:   kategori.MenuUrut,
			MenuItem:   []menu.MenuItem{},
		}
		listKategori = append(listKategori, tempKategory)
	}

	// filter menu berdasarkan role yang login
	for _, item := range menuData {
		allowedRoles := strings.Split(item.LevelUser, "|")
		for _, roleMenu := range allowedRoles {
			menu := menu.MenuItem{
				MenuID:     item.MenuID,
				Kategori:   item.Kategori,
				MenuParent: item.MenuParent,
				MenuIcon:   item.MenuIcon,
				MenuTitle:  item.MenuTitle,
				MenuLink:   item.MenuLink,
				MenuActive: item.MenuActive,
				LevelUser:  item.LevelUser,
				Submenu:    nil,
			}

			if role == roleMenu || roleMenu == "ALL" {
				menuFilter = append(menuFilter, menu)
				break
			}
		}
	}

	// olah data ke dalam submenu
	for _, data := range listKategori {
		resp := menu.ResponseMenu{
			Kategori: data.Kategori,
			Menu:     []menu.MenuItem{},
		}

		var mainMenu menu.MenuItem
		for _, item := range menuFilter {
			if data.KategoriID == item.Kategori {
				mainMenu = item
				for _, dataMenu := range menuFilter {
					if mainMenu.MenuID == dataMenu.MenuParent {
						datatemp := menu.MenuItem{
							MenuID:     dataMenu.MenuID,
							Kategori:   dataMenu.Kategori,
							MenuParent: dataMenu.MenuParent,
							MenuIcon:   dataMenu.MenuIcon,
							MenuTitle:  dataMenu.MenuTitle,
							MenuLink:   dataMenu.MenuLink,
							MenuActive: dataMenu.MenuActive,
							LevelUser:  dataMenu.LevelUser,
							Submenu:    nil,
						}

						for _, data := range menuFilter {
							if datatemp.MenuID == data.MenuParent {
								childMenu := menu.MenuItem{
									MenuID:     data.MenuID,
									Kategori:   data.Kategori,
									MenuParent: data.MenuParent,
									MenuIcon:   data.MenuIcon,
									MenuTitle:  data.MenuTitle,
									MenuLink:   data.MenuLink,
									MenuActive: data.MenuActive,
									LevelUser:  data.LevelUser,
									Submenu:    nil,
								}

								datatemp.Submenu = append(datatemp.Submenu, &childMenu)
							}
						}

						mainMenu.Submenu = append(mainMenu.Submenu, &datatemp)
					}
				}

				resp.Menu = append(resp.Menu, mainMenu)
			}
		}

		if len(resp.Menu) > 0 {
			responseMenu = append(responseMenu, resp)
		}
	}

	return responseMenu, nil
}
