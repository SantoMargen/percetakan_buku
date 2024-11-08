package upload

type Handler struct {
	uploadFileUC uploadFileUC
}

func New(uploadFileUC uploadFileUC) *Handler {
	return &Handler{uploadFileUC: uploadFileUC}
}
