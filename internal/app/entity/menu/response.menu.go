package menu

type MenuItem struct {
	MenuID     string      `json:"menu_id"`
	Kategori   int         `json:"kategori"`
	MenuParent string      `json:"menu_parent"`
	MenuIcon   string      `json:"menu_icon"`
	MenuTitle  string      `json:"menu_title"`
	MenuLink   string      `json:"menu_link"`
	MenuActive int         `json:"menu_active"`
	LevelUser  string      `json:"level_user"`
	Submenu    []*MenuItem `json:"submenu"`
}

type ResponseMenu struct {
	Kategori string     `json:"kategori"`
	Menu     []MenuItem `json:"menu"`
}

type KategoryData struct {
	KategoriID int        `json:"kategori_id"`
	Kategori   string     `json:"kategori"`
	MenuUrut   int        `json:"menu_urut"`
	MenuItem   []MenuItem `json:"menu"`
}

type Kategory struct {
	KategoriID int    `json:"kategori_id"`
	Kategori   string `json:"kategori"`
	MenuUrut   int    `json:"menu_urut"`
}

type Menu struct {
	MenuID     string `json:"menu_id"`
	Kategori   int    `json:"kategori"`
	MenuParent string `json:"menu_parent"`
	MenuIcon   string `json:"menu_icon"`
	MenuTitle  string `json:"menu_title"`
	MenuLink   string `json:"menu_link"`
	MenuActive int    `json:"menu_active"`
	LevelUser  string `json:"level_user"`
}
