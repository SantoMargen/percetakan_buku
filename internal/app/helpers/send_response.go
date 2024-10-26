package helpers

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := entity.Response[interface{}]{
		Data: map[string]interface{}{
			"message":      message,
			"responseCode": statusCode,
			"status":       "Success",
			"data":         data,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, statusCode int, errType, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResponse := entity.Response[map[string]interface{}]{
		Data: map[string]interface{}{
			"message":      errMessage,
			"errorType":    errType,
			"responseCode": statusCode,
			"data":         nil,
		},
	}

	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
