package main

import (
	"log"
	"net/http"
	_ "test_project/test/docs"
	"test_project/test/internal/config"
	"test_project/test/internal/handler"
	"test_project/test/internal/middleware"
	rt "test_project/test/internal/router"
	"test_project/test/internal/service"
	"test_project/test/internal/storage"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Ideas API
// @version 1.1
// @description API for managing project ideas
// @host go-ideas-api.onrender.com
// @BasePath /v1
// @schemes https
func main() {
	// store := storage.NewJsonStore("data/ideas.json")
	_ = godotenv.Load(".env")
	dbconfig := config.NewDBConfig()

	pg, err := storage.NewPostgresStore(dbconfig.GetDSNPG())
	if err != nil {
		log.Fatalf("Failed to initialize postgres db: %v", err)
	}
	store := pg

	s := service.NewIdeaService(store)
	authService := service.NewUserService(store)

	h := handler.NewIdeaHandler(s)
	authHandler := handler.NewAuthHandler(authService)

	router := http.NewServeMux()
	v1Routes := rt.SetupRoutes(h, authHandler)

	router.Handle("/v1/", http.StripPrefix("/v1", middleware.CORS(middleware.Logging(v1Routes))))
	router.Handle("/", middleware.CORS(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server starting at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
