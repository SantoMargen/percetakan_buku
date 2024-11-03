package routes

import (
	handlerCategory "siap_app/internal/app/handler/category"
	handlerLevelUser "siap_app/internal/app/handler/level_users"
	handlerMenu "siap_app/internal/app/handler/menu"
	handlerNotification "siap_app/internal/app/handler/notification"
	handlerPaper "siap_app/internal/app/handler/papers"
	handlerPublisher "siap_app/internal/app/handler/publishers"
	handlerUser "siap_app/internal/app/handler/user"

	"siap_app/internal/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(
	r chi.Router,
	userHandler *handlerUser.Handler,
	menuHandler *handlerMenu.Handler,
	handlerLevelUser *handlerLevelUser.Handler,
	publisherHandler *handlerPublisher.Handler,
	handlerCategory *handlerCategory.Handler,
	handlerPaper *handlerPaper.Handler,
	handlerNotif *handlerNotification.Handler,

) {
	SetUserRoutes(r, userHandler)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthorizationMiddleware)
		SetMenuRoutes(r, menuHandler)
		SetLevelUserRoutes(r, handlerLevelUser)
		SetPublisherRoutes(r, publisherHandler)
		SetCategoryRoutes(r, handlerCategory)
		SetPaperRoutes(r, handlerPaper)
		SetNotificationRoutes(r, handlerNotif)
	})
}
