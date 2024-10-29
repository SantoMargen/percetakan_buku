package user

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.RegisterRequest
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

	err = h.userUC.CreateUser(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "User created successfully", http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.LoginRequest
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

	resp, err := h.userUC.LoginUser(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, resp, "login user successfully", http.StatusOK)
}
