package routes

import (
	notificationHandler "siap_app/internal/app/handler/notification"

	"github.com/go-chi/chi/v5"
)

func SetNotificationRoutes(r chi.Router, h *notificationHandler.Handler) {
	r.Route("/notification", func(r chi.Router) {
		r.Post("/add", h.CreateLogNotif)
		r.Post("/delete", h.DeleteNotification)
		r.Post("/all", h.GetNotificationAll)
	})
}
