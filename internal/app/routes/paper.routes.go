package routes

import (
	handleCategory "siap_app/internal/app/handler/papers"

	"github.com/go-chi/chi/v5"
)

func SetPaperRoutes(r chi.Router, h *handleCategory.Handler) {
	r.Route("/paper", func(r chi.Router) {
		r.Post("/add", h.CreatePaper)
		r.Post("/delete", h.DeletePaper)
		r.Post("/by_id", h.GetPaperById)
		r.Post("/update", h.UpdatePaper)
	})
}
