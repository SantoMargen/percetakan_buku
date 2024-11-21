package routes

import (
	handleCategory "siap_app/internal/app/handler/papers"

	"github.com/go-chi/chi/v5"
)

func SetPaperRoutes(r chi.Router, h *handleCategory.Handler) {
	r.Route("/paper", func(r chi.Router) {
		r.Post("/add", h.CreatePaper)
		r.Post("/delete", h.DeletePaper)
		r.Post("/by-id", h.GetPaperById)
		r.Post("/by-user-id", h.GetDetailPaperUserById)
		r.Post("/update", h.UpdatePaper)
		r.Post("/assign-paper", h.AssignPaper)
		r.Post("/assign-task-publisher", h.AssignPaperPublisher)
		r.Post("/approval-paper", h.ApprovalPaper)
	})
}
