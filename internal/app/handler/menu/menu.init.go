package menu

type Handler struct {
	menuUC menuUC
}

func New(menuUC menuUC) *Handler {
	return &Handler{menuUC: menuUC}
}
