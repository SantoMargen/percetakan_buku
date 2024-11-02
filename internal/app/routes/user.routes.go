package routes

import (
	handlerUser "siap_app/internal/app/handler/user"
	"siap_app/internal/app/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

func SetUserRoutes(r chi.Router, h *handlerUser.Handler, redis *redis.Client) {
	r.Post("/login", h.LoginUser)
	r.Post("/register", h.CreateUser)

	r.Group(func(protected chi.Router) {
		protected.Use(middlewares.AuthorizationMiddleware(redis))
		protected.Post("/logout", h.Logout)
		protected.Route("/admin", func(admin chi.Router) {
			admin.Use(middlewares.AdminMiddleware)
			admin.Post("/create", h.CreateUserByAdmin)
		})
	})
}
