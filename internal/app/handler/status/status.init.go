package status

type Handler struct {
	statusUC statusUC
}

func New(statusUC statusUC) *Handler {
	return &Handler{statusUC: statusUC}
}
