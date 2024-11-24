package routes

import (
	handleCategory "siap_app/internal/app/handler/status"

	"github.com/go-chi/chi/v5"
)

func SetStatusRoutes(r chi.Router, h *handleCategory.Handler) {
	r.Route("/status", func(r chi.Router) {
		r.Post("/all", h.GetStatusAll)
	})
}
