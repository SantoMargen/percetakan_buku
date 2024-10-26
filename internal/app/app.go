package app

import (
	"siap_app/config"
	userHandler "siap_app/internal/app/handler/user"
	"siap_app/internal/app/middlewares"
	userRepo "siap_app/internal/app/repository/user"
	"siap_app/internal/app/routes"
	userUC "siap_app/internal/app/usecase/user"

	"log"
	"net/http"

	"siap_app/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type App struct {
	Router *chi.Mux
	DB     *sqlx.DB
}

func NewApp() *App {
	app := &App{}
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize dependencies
	dbConfig := config.LoadDBConfig()
	app.DB, err = db.InitDB(*dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = chi.NewRouter()
	app.Router.Use(middlewares.CORSMiddleware)
	app.Router.Use(middlewares.LoggingMiddleware)

	userRepository, err := userRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize user repository: %v", err)

	}
	logrus.Info("Init User repository")

	userUC := userUC.New(userRepository)
	logrus.Info("Init user usecase")

	userHandler := userHandler.New(userUC)
	logrus.Info("Init user handler")

	// // Register user routes
	// userHandler.Routes(app.Router)
	// evenHandler.Routes(app.Router)
	// ticketHandler.Routes(app.Router)

	// routes.SetupRoutes(app.Router, userHandler, evenHandler, ticketHandler)
	routes.SetupRoutes(app.Router, userHandler)

	return app
}

func (a *App) Run(addr string) error {
	logrus.Info("Starting Server...")
	logrus.Info("App running on port:8080")
	return http.ListenAndServe(addr, a.Router)
}
