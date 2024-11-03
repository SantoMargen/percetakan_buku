package notification

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity"
	generalResponse "siap_app/internal/app/entity"
	"siap_app/internal/app/entity/notification"
	"siap_app/internal/app/helpers"
)

func (h *Handler) GetNotificationAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input notification.NotificationPagination

	roleID, ok := r.Context().Value(entity.RoleKey).(string)
	if !ok || roleID == "" {
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

	resp, total, err := h.notificationUC.GetNotificationAll(ctx, input, roleID)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	responseData := generalResponse.ResponsePagination{
		Total: total,
		Data:  resp,
	}

	helpers.SendSuccessResponse(w, responseData, "Notification fetch all successfully", http.StatusOK)
}

func (h *Handler) CreateLogNotif(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input notification.RequestNotification

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

	sentSenderNotification := notification.SentRequestNotification{
		RequestNotification: input,
		Sender:              userId,
	}

	err = h.notificationUC.CreateLogNotif(ctx, sentSenderNotification)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Create notification succesfully", http.StatusCreated)
}

func (h *Handler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input notification.RequestNotificationById

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

	err = h.notificationUC.DeleteNotification(ctx, input.ID, userId)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Delete notification successfully", http.StatusCreated)
}
