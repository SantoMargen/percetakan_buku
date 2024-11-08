package routes

import (
	handlerUpload "siap_app/internal/app/handler/upload"

	"github.com/go-chi/chi/v5"
)

func SentUploadRoutes(r chi.Router, h *handlerUpload.Handler) {
	r.Route("/upload", func(r chi.Router) {
		r.Post("/sent-file", h.UploadFile)
	})
}
