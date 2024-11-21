package levelusers

type Handler struct {
	levelUserUC levelUserUC
}

func New(levelUserUC levelUserUC) *Handler {
	return &Handler{levelUserUC: levelUserUC}
}
