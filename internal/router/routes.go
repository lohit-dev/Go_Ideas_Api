package router

import (
	"net/http"
	"test_project/test/internal/handler"
	"test_project/test/internal/middleware"
)

func SetupRoutes(ideaHandler *handler.IdeaHandler, authHandler *handler.AuthHandler, voteHandler *handler.VoteHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Auth
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/user", authHandler.DeleteUser)
	mux.HandleFunc("GET /auth/users", authHandler.GetAllUsers)
	mux.HandleFunc("GET /auth/user/{username}", authHandler.GetUserByUsername)

	// Idea
	mux.Handle("POST /idea", middleware.Auth(http.HandlerFunc(ideaHandler.CreateIdea)))
	mux.Handle("GET /idea/{id}", http.HandlerFunc(ideaHandler.GetIdea))
	mux.Handle("GET /ideas", http.HandlerFunc(ideaHandler.GetAllIdeas))
	mux.Handle("POST /idea/{id}", middleware.Auth(http.HandlerFunc(ideaHandler.UpdateIdea)))
	mux.Handle("DELETE /idea/{id}", middleware.Auth(http.HandlerFunc(ideaHandler.DeleteIdea)))

	// Voting
	mux.Handle("POST /idea/{id}/vote", (http.HandlerFunc(voteHandler.AddVote)))
	mux.Handle("DELETE /idea/{id}/vote", (http.HandlerFunc(voteHandler.RemoveVote)))
	mux.Handle("GET /idea/{id}/vote/status", http.HandlerFunc(voteHandler.HasUserVoted))
	mux.Handle("GET /idea/{id}/votes", http.HandlerFunc(voteHandler.GetVoteCount))

	return mux
}
