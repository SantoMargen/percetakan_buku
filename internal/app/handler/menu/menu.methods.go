package menu

import (
	"fmt"
	"net/http"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/helpers"
)

func (h *Handler) GetMenu(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	role, ok := r.Context().Value(entity.RoleKey).(string)
	if !ok || role == "" {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	fmt.Println(role, "================== ROLE")

	resp, err := h.menuUC.GetMenu(ctx, role)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, resp, "login user successfully", http.StatusOK)
}
