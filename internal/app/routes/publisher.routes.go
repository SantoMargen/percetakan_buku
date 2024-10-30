package routes

import (
	handlePublisher "siap_app/internal/app/handler/publishers"

	"github.com/go-chi/chi/v5"
)

func SetPublisherRoutes(r chi.Router, h *handlePublisher.Handler) {
	r.Route("/publisher", func(r chi.Router) {
		r.Post("/add", h.CreatePublisher)
		r.Post("/update", h.UpdatePublisher)
		r.Post("/delete", h.DeletePublisher)
		r.Post("/by_id", h.GetPublisherById)
	})
}
