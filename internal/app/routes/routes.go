package routes

import (
	handlerCategory "siap_app/internal/app/handler/category"
	handlerLevelUser "siap_app/internal/app/handler/level_users"
	handlerMenu "siap_app/internal/app/handler/menu"
	handlerPublisher "siap_app/internal/app/handler/publishers"
	handlerUser "siap_app/internal/app/handler/user"

	"siap_app/internal/app/middlewares"

	"github.com/go-chi/chi/v5"
)

// func SetupRoutes(r chi.Router, userHandler *handlerUser.Handler, menuHandler *handlerMenu.Handler) {
// 	SetUserRoutes(r, userHandler)
// 	SetMenuRoutes(r, menuHandler)
// }

func SetupRoutes(
	r chi.Router,
	userHandler *handlerUser.Handler,
	menuHandler *handlerMenu.Handler,
	handlerLevelUser *handlerLevelUser.Handler,
	publisherHandler *handlerPublisher.Handler,
	handlerCategory *handlerCategory.Handler,
) {
	SetUserRoutes(r, userHandler)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthorizationMiddleware)
		SetMenuRoutes(r, menuHandler)
		SetLevelUserRoutes(r, handlerLevelUser)
		SetPublisherRoutes(r, publisherHandler)
		SetCategoryRoutes(r, handlerCategory)
	})
}
