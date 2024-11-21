package papers

type Handler struct {
	paperUC paperUC
}

func New(paperUC paperUC) *Handler {
	return &Handler{paperUC: paperUC}
}
