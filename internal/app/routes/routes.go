package routes

import (
	handlerUser "siap_app/internal/app/handler/user"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router, userHandler *handlerUser.Handler) {
	SetUserRoutes(r, userHandler)
	// SetEventRoutes(r, eventHandler)
	// SetTicketRoutes(r, ticketHandler)
}
