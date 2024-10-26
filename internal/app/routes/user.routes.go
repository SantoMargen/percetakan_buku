package routes

import (
	handlerUser "siap_app/internal/app/handler/user"
	"siap_app/internal/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func SetUserRoutes(r chi.Router, h *handlerUser.Handler) {
	r.Post("/login", h.LoginUser)
	// r.Post("/logout", h.LogoutUser)
	r.Route("/users", func(r chi.Router) {
		r.Use(middlewares.AuthorizationMiddleware)
		// r.Get("/", h.GetUsers)
		r.Post("/register", h.CreateUser)
	})
}
