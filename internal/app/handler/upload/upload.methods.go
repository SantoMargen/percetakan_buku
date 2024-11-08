package upload

import (
	"net/http"
	"siap_app/internal/app/helpers"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	path := r.FormValue("path")

	if path == "" {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", " PATH NOT BE EMPTY")
	}
	resFile, err := helpers.UploadFileHandler(ctx, r, path)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	errUploadFile := h.uploadFileUC.UploadFile(ctx, resFile)

	if errUploadFile != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", errUploadFile.Error())
		return
	}

	helpers.SendSuccessResponse(w, resFile, "Upload file successfully", http.StatusCreated)
}
