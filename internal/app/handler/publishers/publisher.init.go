package publishers

type Handler struct {
	publisherUC publisherUC
}

func New(publisherUC publisherUC) *Handler {
	return &Handler{publisherUC: publisherUC}
}
