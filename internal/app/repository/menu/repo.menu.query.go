package menu

const (
	queryGetKategoryMenu string = `
		SELECT 
			kategori_id,
			kategori,
			menu_urut   
		FROM kategory_menu ORDER BY menu_urut ASC`

	queryGetMenu string = `
	SELECT 
		menu_id, 
		menu_parent,
		kategori,
		menu_icon,
		menu_title,
		menu_link, 
		menu_active, 
		level_user 
	FROM menu ORDER BY menu_id ASC`
)
