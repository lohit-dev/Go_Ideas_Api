package main

import (
	"log"
	"net/http"
	_ "test_project/test/docs"
	"test_project/test/internal/handler"
	rt "test_project/test/internal/router"
	"test_project/test/internal/service"
	"test_project/test/internal/storage"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Ideas API
// @version 1.0
// @description API for managing project ideas
// @host localhost:8080
// @BasePath /v1
// @schemes http
func main() {
	store := storage.NewJsonStore("data/ideas.json")
	service := service.NewIdeaService(store)
	h := handler.NewIdeaHandler(service)

	router := http.NewServeMux()
	v1Routes := rt.SetupRoutes(h)
	router.Handle("/v1/", http.StripPrefix("/v1", v1Routes))

	router.Handle("/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server starting at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
