package router

import (
	"net/http"
	"test_project/test/internal/handler"
	"test_project/test/internal/middleware"
)

func SetupRoutes(ideaHandler *handler.IdeaHandler, authHandler *handler.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Auth
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/register", authHandler.Register)

	// Idea
	mux.Handle("POST /idea", middleware.Auth(http.HandlerFunc(ideaHandler.CreateIdea)))
	mux.Handle("GET /idea/{id}", middleware.Auth(http.HandlerFunc(ideaHandler.GetIdea)))
	mux.Handle("GET /idea", middleware.Auth(http.HandlerFunc(ideaHandler.GetAllIdeas)))
	mux.Handle("POST /idea/{id}", middleware.Auth(http.HandlerFunc(ideaHandler.UpdateIdea)))
	mux.Handle("DELETE /idea/{id}", middleware.Auth(http.HandlerFunc(ideaHandler.DeleteIdea)))

	return mux
}
