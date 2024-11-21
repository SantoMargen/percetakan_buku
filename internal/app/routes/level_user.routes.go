package routes

import (
	handlerLevelUser "siap_app/internal/app/handler/level_users"

	"github.com/go-chi/chi/v5"
)

func SetLevelUserRoutes(r chi.Router, h *handlerLevelUser.Handler) {
	r.Route("/level_user", func(r chi.Router) {
		r.Get("/list", h.GetLevelUsers)
		r.Post("/byid", h.GetLevelUserBYID)
	})
}
