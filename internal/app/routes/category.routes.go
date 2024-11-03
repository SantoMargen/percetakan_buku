package routes

import (
	handleCategory "siap_app/internal/app/handler/category"

	"github.com/go-chi/chi/v5"
)

func SetCategoryRoutes(r chi.Router, h *handleCategory.Handler) {
	r.Route("/category", func(r chi.Router) {
		r.Post("/add", h.CreateCategory)
		r.Post("/update", h.UpdateCategory)
		r.Post("/delete", h.DeleteCategory)
		r.Post("/by_id", h.GetCategoryById)
		r.Post("/all", h.GetCategoryAll)
	})
}
