package papers

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/entity/papers"
	"siap_app/internal/app/helpers"
)

func (h *Handler) CreatePaper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input papers.RequestPaper

	userId, ok := r.Context().Value(entity.UserIDKey).(int)
	if !ok || userId == 0 {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.paperUC.CreatePaper(ctx, input, userId)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Submission paper successfully", http.StatusCreated)
}

func (h *Handler) DeletePaper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input papers.RequestPaperById

	userId, ok := r.Context().Value(entity.UserIDKey).(int)
	if !ok || userId == 0 {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.paperUC.DeletePaper(ctx, input.ID, userId)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Delete submission paper successfully", http.StatusCreated)
}

func (h *Handler) GetPaperById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input papers.RequestPaperById
	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failed umarshal data", err.Error())
		return
	}

	data, err := h.paperUC.GetPaperById(ctx, input.ID)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, data, "Success get retrieve paper", http.StatusOK)
}

func (h *Handler) UpdatePaper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input papers.RequestPaperUpdate

	userId, ok := r.Context().Value(entity.UserIDKey).(int)
	if !ok || userId == 0 {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.paperUC.UpdatePaper(ctx, input, userId)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Submission update successfully", http.StatusCreated)
}
