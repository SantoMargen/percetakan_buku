package upload

import (
	"net/http"
	"siap_app/internal/app/entity/upload"
	"siap_app/internal/app/helpers"
	"strconv"
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

	var getResFileUpload []upload.ResponseUpload

	for i, value := range resFile {
		getFile := strconv.Itoa(i)
		getFileUpload := upload.ResponseUpload{
			IDFile:   value.IDFile + getFile,
			Filename: value.Filename,
			Filetype: value.Filetype,
		}

		getResFileUpload = append(getResFileUpload, getFileUpload)

	}

	helpers.SendSuccessResponse(w, getResFileUpload, "Upload file successfully", http.StatusCreated)
}
