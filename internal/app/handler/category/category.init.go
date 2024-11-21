package category

type Handler struct {
	categoryUC categoryUC
}

func New(categoryUC categoryUC) *Handler {
	return &Handler{categoryUC: categoryUC}
}
