package routes

import (
	handlerUser "siap_app/internal/app/handler/user"

	"github.com/go-chi/chi/v5"
)

func SetUserRoutes(r chi.Router, h *handlerUser.Handler) {
	r.Post("/login", h.LoginUser)
	r.Post("/register", h.CreateUser)
}
