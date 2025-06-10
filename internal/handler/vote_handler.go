package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_project/test/internal/service"
	utils "test_project/test/pkg"

	"github.com/google/uuid"
)

type VoteHandler struct {
	service *service.VoteService
}

func NewVoteHandler(service *service.VoteService) *VoteHandler {
	return &VoteHandler{service}
}

// AddVote godoc
// @Summary Add a vote to an idea
// @Description Allows a user to vote for a specific idea
// @Tags Votes
// @Accept json
// @Produce json
// @Param id path string true "Idea ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Vote added successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid ID format"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Failure 409 {object} map[string]string "Conflict - user has already voted"
// @Failure 500 {object} map[string]string "Server error"
// @Router /idea/{id}/vote [post]
func (h *VoteHandler) AddVote(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ideaIdStr := r.PathValue("id")
	if ideaIdStr == "" {
		http.Error(w, "missing idea id parameter", http.StatusBadRequest)
		return
	}

	ideaID, err := uuid.Parse(ideaIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid idea ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.AddVote(userID, ideaID)
	if result.Err != nil {
		// Check if it's a "already voted" error
		if result.Err.Error() == "User has already voted" {
			http.Error(w, result.Err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("failed to add vote: %v", result.Err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": result.Data}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// RemoveVote godoc
// @Summary Remove a vote from an idea
// @Description Allows a user to remove their vote from a specific idea
// @Tags Votes
// @Accept json
// @Produce json
// @Param id path string true "Idea ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string "Vote removed successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid ID format"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]string "Not found - user has not voted for this idea"
// @Failure 500 {object} map[string]string "Server error"
// @Router /idea/{id}/vote [delete]
func (h *VoteHandler) RemoveVote(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ideaIdStr := r.PathValue("id")
	if ideaIdStr == "" {
		http.Error(w, "missing idea id parameter", http.StatusBadRequest)
		return
	}

	ideaID, err := uuid.Parse(ideaIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid idea ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.RemoveVote(userID, ideaID)
	if result.Err != nil {
		// Check if it's a "not voted" error
		if result.Err.Error() == "User has not voted for this idea" {
			http.Error(w, result.Err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("failed to remove vote: %v", result.Err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": result.Data}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// HasUserVoted godoc
// @Summary Check if user has voted for an idea
// @Description Checks whether the authenticated user has voted for a specific idea
// @Tags Votes
// @Produce json
// @Param id path string true "Idea ID"
// @Security BearerAuth
// @Success 200 {object} map[string]bool "Returns voting status"
// @Failure 400 {object} map[string]string "Bad request - invalid ID format"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Failure 500 {object} map[string]string "Server error"
// @Router /idea/{id}/vote/status [get]
func (h *VoteHandler) HasUserVoted(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ideaIdStr := r.PathValue("id")
	if ideaIdStr == "" {
		http.Error(w, "missing idea id parameter", http.StatusBadRequest)
		return
	}

	ideaID, err := uuid.Parse(ideaIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid idea ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.HasUserVoted(userID, ideaID)
	if result.Err != nil {
		http.Error(w, fmt.Sprintf("failed to check vote status: %v", result.Err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]bool{"hasVoted": result.Data}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// GetVoteCount godoc
// @Summary Get vote count for an idea
// @Description Retrieves the total number of votes for a specific idea
// @Tags Votes
// @Produce json
// @Param id path string true "Idea ID"
// @Success 200 {object} map[string]int "Returns vote count"
// @Failure 400 {object} map[string]string "Bad request - invalid ID format"
// @Failure 500 {object} map[string]string "Server error"
// @Router /idea/{id}/votes [get]
func (h *VoteHandler) GetVoteCount(w http.ResponseWriter, r *http.Request) {
	ideaIdStr := r.PathValue("id")
	if ideaIdStr == "" {
		http.Error(w, "missing idea id parameter", http.StatusBadRequest)
		return
	}

	ideaID, err := uuid.Parse(ideaIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid idea ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.GetVoteCount(ideaID)
	if result.Err != nil {
		http.Error(w, fmt.Sprintf("failed to get vote count: %v", result.Err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]int{"voteCount": result.Data}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
