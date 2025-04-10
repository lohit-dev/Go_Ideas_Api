package router

import (
	"net/http"
	"test_project/test/internal/handler"
)

func SetupRoutes(ideaHandler *handler.IdeaHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /idea", ideaHandler.CreateIdea)
	mux.HandleFunc("GET /idea/{id}", ideaHandler.GetIdea)
	mux.HandleFunc("GET /idea", ideaHandler.GetAllIdeas)
	mux.HandleFunc("POST /idea/{id}", ideaHandler.UpdateIdea)
	mux.HandleFunc("DELETE /idea/{id}", ideaHandler.DeleteIdea)

	return mux
}
