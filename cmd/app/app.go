package app

import (
	"net/http"
	"test_project/test/internal/config"
	"test_project/test/internal/handler"
	"test_project/test/internal/middleware"
	"test_project/test/internal/service"
	"test_project/test/internal/storage"

	rt "test_project/test/internal/router"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	server *http.Server
}

func NewApp() (*App, error) {
	// Load environment variables
	_ = godotenv.Load(".env")

	// Initialize storage
	store, err := initStorage()
	if err != nil {
		return nil, err
	}

	// Initialize
	services := initServices(store)
	handlers := initHandlers(services)
	router := setupRouter(handlers)

	return &App{
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}, nil
}

func (a *App) Start() error {
	return a.server.ListenAndServe()
}

func initStorage() (*storage.PostgresStore, error) {
	dbconfig := config.NewDBConfig()
	pg, err := storage.NewPostgresStore(dbconfig.GetDSNPG())
	if err != nil {
		return nil, err
	}
	return pg, nil
}

type Services struct {
	IdeaService *service.IdeaService
	UserService *service.UserService
	VoteService *service.VoteService
}

func initServices(store *storage.PostgresStore) *Services {
	return &Services{
		IdeaService: service.NewIdeaService(store),
		UserService: service.NewUserService(store),
		VoteService: service.NewVoteService(store),
	}
}

type Handlers struct {
	IdeaHandler *handler.IdeaHandler
	AuthHandler *handler.AuthHandler
	VoteHandler *handler.VoteHandler
}

func initHandlers(services *Services) *Handlers {
	return &Handlers{
		IdeaHandler: handler.NewIdeaHandler(services.IdeaService),
		AuthHandler: handler.NewAuthHandler(services.UserService),
		VoteHandler: handler.NewVoteHandler(services.VoteService),
	}
}

func setupRouter(handlers *Handlers) *http.ServeMux {
	router := http.NewServeMux()

	// API routes
	v1Routes := rt.SetupRoutes(handlers.IdeaHandler, handlers.AuthHandler, handlers.VoteHandler)
	router.Handle("/v1/", http.StripPrefix("/v1", middleware.CORS(middleware.Logging(v1Routes))))

	// Swagger documentation
	router.Handle("/", middleware.CORS(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)))

	return router
}
