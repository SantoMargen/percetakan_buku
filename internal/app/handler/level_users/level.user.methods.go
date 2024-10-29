package levelusers

import (
	"encoding/json"
	"net/http"
	levelusers "siap_app/internal/app/entity/level_users"
	"siap_app/internal/app/helpers"
)

func (h *Handler) GetLevelUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.levelUserUC.GetLevelUsers(ctx)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, data, "Success get retrieve level users", http.StatusOK)
}

func (h *Handler) GetLevelUserBYID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input levelusers.RequestLevelUserByID
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

	data, err := h.levelUserUC.GetLevelUsersByID(ctx, input.ID)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, data, "Success get retrieve level user", http.StatusOK)
}
