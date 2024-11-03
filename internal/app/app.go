package app

import (
	"siap_app/config"
	categoryHandler "siap_app/internal/app/handler/category"
	levelUserHandler "siap_app/internal/app/handler/level_users"
	menuHandler "siap_app/internal/app/handler/menu"
	notificationHandler "siap_app/internal/app/handler/notification"
	paperHandler "siap_app/internal/app/handler/papers"
	publisherHandler "siap_app/internal/app/handler/publishers"
	userHandler "siap_app/internal/app/handler/user"

	"siap_app/internal/app/middlewares"
	categoryRepo "siap_app/internal/app/repository/category"
	levelUserRepo "siap_app/internal/app/repository/level_users"
	logLoginRepo "siap_app/internal/app/repository/log_login"
	menuRepo "siap_app/internal/app/repository/menu"
	notificationRepo "siap_app/internal/app/repository/notification"
	paperRepo "siap_app/internal/app/repository/paper"
	publisherRepo "siap_app/internal/app/repository/publishers"
	redisRepo "siap_app/internal/app/repository/redis"
	userRepo "siap_app/internal/app/repository/user"

	"siap_app/internal/app/routes"
	categoryUC "siap_app/internal/app/usecase/category"
	levelUserUC "siap_app/internal/app/usecase/level_users"
	menuUC "siap_app/internal/app/usecase/menu"
	notificationUC "siap_app/internal/app/usecase/notification"
	paperUC "siap_app/internal/app/usecase/papers"
	publisherUC "siap_app/internal/app/usecase/publisher"
	userUC "siap_app/internal/app/usecase/user"

	"log"
	"net/http"

	"siap_app/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type App struct {
	Router *chi.Mux
	DB     *sqlx.DB
	Redis  *redis.Client
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

	app.Redis, err = db.RedisInit(*dbConfig)
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

	menuRepository, err := menuRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize user repository: %v", err)
	}
	logrus.Info("Init Menu repository")

	levelUserRepository, err := levelUserRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize level user repository: %v", err)
	}
	logrus.Info("Init Level User repository")

	redisRepository, err := redisRepo.New(app.Redis)
	if err != nil {
		logrus.Fatalf("Failed to initialize redis repository: %v", err)
	}
	logrus.Info("Init redis repository")

	logLoginRepository, err := logLoginRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize log login repository: %v", err)
	}
	logrus.Info("Init log login repository")

	publisherRepository, err := publisherRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize publisher repository: %v", err)
	}
	logrus.Info("Init  Publisher repository")

	categoryRepository, err := categoryRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize publisher repository: %v", err)
	}
	logrus.Info("Init Category repository")

	paperRepository, err := paperRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize paper repository: %v", err)
	}
	logrus.Info("Init Paper repository")

	notifRepository, err := notificationRepo.New(app.DB)
	if err != nil {
		logrus.Fatalf("Failed to initialize notification repository: %v", err)
	}
	logrus.Info("Init Paper repository")

	userUC := userUC.New(userRepository, redisRepository, logLoginRepository)
	logrus.Info("Init user usecase")
	menuUC := menuUC.New(menuRepository)
	logrus.Info("Init menu usecase")
	levelUserUC := levelUserUC.New(levelUserRepository)
	logrus.Info("Init level user usecase")
	publisherUC := publisherUC.New(publisherRepository)
	logrus.Info("Init publisher")
	categoryUC := categoryUC.New(categoryRepository)
	logrus.Info("Init publisher")
	paperUC := paperUC.New(
		paperRepository,
		publisherRepository,
		notifRepository,
	)
	logrus.Info("Init publisher")
	notifUC := notificationUC.New(notifRepository)
	logrus.Info("Init notification")

	userHandler := userHandler.New(userUC)
	logrus.Info("Init user handler")
	menuHandler := menuHandler.New(menuUC)
	logrus.Info("Init menu handler")
	lvelUserHandler := levelUserHandler.New(levelUserUC)
	logrus.Info("Init level user handler")
	publisherHandler := publisherHandler.New(publisherUC)
	logrus.Info("Init publisher handler")
	categoryHandler := categoryHandler.New(categoryUC)
	logrus.Info("Init category handler")
	paperHandler := paperHandler.New(paperUC)
	logrus.Info("Init paper handler")
	notifHandler := notificationHandler.New(notifUC)
	logrus.Info("Init paper handler")

	// // Register user routes
	// userHandler.Routes(app.Router)
	// evenHandler.Routes(app.Router)
	// ticketHandler.Routes(app.Router)

	// routes.SetupRoutes(app.Router, userHandler, evenHandler, ticketHandler)
	routes.SetupRoutes(
		app.Router,
		app.Redis,
		userHandler,
		menuHandler,
		lvelUserHandler,
		publisherHandler,
		categoryHandler,
		paperHandler,
		notifHandler,
	)

	return app
}

func (a *App) Run(addr string) error {
	logrus.Info("Starting Server...")
	logrus.Info("App running on port:8080")
	return http.ListenAndServe(addr, a.Router)
}
