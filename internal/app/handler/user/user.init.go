package user

type Handler struct {
	userUC userUC
}

func New(userUC userUC) *Handler {
	return &Handler{userUC: userUC}
}
