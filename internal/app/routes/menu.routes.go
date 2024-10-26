package routes

import (
	handlerMenu "siap_app/internal/app/handler/menu"

	"github.com/go-chi/chi/v5"
)

func SetMenuRoutes(r chi.Router, h *handlerMenu.Handler) {
	r.Route("/menu", func(r chi.Router) {
		r.Get("/list", h.GetMenu)
	})
}
