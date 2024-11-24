package status

import (
	"encoding/json"
	"net/http"
	generalResponse "siap_app/internal/app/entity"
	"siap_app/internal/app/entity/status"
	"siap_app/internal/app/helpers"
)

func (h *Handler) GetStatusAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input status.PaginationStatus

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

	resp, total, err := h.statusUC.GetStatusAll(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	responseData := generalResponse.ResponsePagination{
		Total: total,
		Data:  resp,
	}

	helpers.SendSuccessResponse(w, responseData, "Category fetch all successfully", http.StatusOK)
}
