package publishers

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity/publishers"
	"siap_app/internal/app/helpers"
)

func (h *Handler) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input publishers.PublisherRequest
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

	err = h.publisherUC.CreatePublisher(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Publisher created successfully", http.StatusCreated)
}

func (h *Handler) GetPublisherById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input publishers.RequestPublishersByID
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

	data, err := h.publisherUC.GetPublisherById(ctx, input.ID)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, data, "Success get retrieve publisher", http.StatusOK)
}
func (h *Handler) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input publishers.RequestUpdatePublisher
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

	err = h.publisherUC.UpdatePublisher(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Publisher update successfully", http.StatusCreated)
}

func (h *Handler) DeletePublisher(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input publishers.RequestPublishersByID
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

	err = h.publisherUC.DeletePublisher(ctx, input.ID)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Publisher delete successfully", http.StatusCreated)
}
