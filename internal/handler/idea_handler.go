package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_project/test/internal/model"
	"test_project/test/internal/service"
	utils "test_project/test/pkg"
	"time"

	"github.com/google/uuid"
)

type IdeaHandler struct {
	service *service.IdeaService
}

func NewIdeaHandler(service *service.IdeaService) *IdeaHandler {
	return &IdeaHandler{service}
}

// CreateIdea godoc
// @Summary Create a new idea
// @Description Creates a new idea in the system with the provided details
// @Tags Ideas
// @Accept json
// @Produce json
// @Param idea body model.CreateIdeaPayload true "Idea object with title, description, tech stack, and tags"
// @Success 201 {object} map[string]string "Returns a success message with the created idea ID"
// @Failure 400 {object} map[string]string "Bad request - invalid payload format or missing required fields"
// @Failure 500 {object} map[string]string "Server error - database or internal processing error"
// @Router /idea [post]
func (h *IdeaHandler) CreateIdea(w http.ResponseWriter, r *http.Request) {
	var createPayload model.CreateIdeaPayload

	if err := json.NewDecoder(r.Body).Decode(&createPayload); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode idea: %v", err), http.StatusBadRequest)
		return
	}

	idea := model.Idea{
		ID:          uuid.MustParse(utils.GenId()),
		Title:       createPayload.Title,
		Description: createPayload.Description,
		TechStack:   createPayload.TechStack,
		Tags:        createPayload.Tags,
		Status:      createPayload.Status,
		Votes:       []model.Vote{},
		RequestedBy: "anonymous",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if idea.Status == "" {
		idea.Status = model.Requested
	}

	result := h.service.CreateIdea(idea)
	if result.Err != nil {
		http.Error(w, fmt.Sprintf("failed to create idea: %v", result.Err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"result": result.Data}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// GetAllIdeas godoc
// @Summary Get all ideas
// @Description Retrieves all ideas from the system
// @Tags Ideas
// @Produce json
// @Success 200 {array} model.Idea
// @Failure 500 {object} error "Server error"
// @Router /idea [get]
func (h *IdeaHandler) GetAllIdeas(w http.ResponseWriter, r *http.Request) {
	result := h.service.GetAllIdeas()
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Data)
}

// GetIdea godoc
// @Summary Get a specific idea by ID
// @Description Retrieves a single idea by its ID
// @Tags Ideas
// @Produce json
// @Param id path string true "Idea ID"
// @Success 200 {object} model.Idea
// @Failure 400 {object} error "Invalid ID format"
// @Failure 404 {object} error "Idea not found"
// @Router /idea/{id} [get]
func (h *IdeaHandler) GetIdea(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.GetIdea(id)
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Data)
}

// UpdateIdea godoc
// @Summary Update an existing idea
// @Description Updates an idea's information in the system
// @Tags Ideas
// @Accept json
// @Produce json
// @Param id path string true "Idea ID"
// @Param idea body model.UpdateIdeaPayload true "Updated idea object"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} error "Invalid request or ID format"
// @Failure 404 {object} error "Idea not found"
// @Failure 500 {object} error "Server error"
// @Router /idea/{id} [post]
func (h *IdeaHandler) UpdateIdea(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	var updatePayload model.UpdateIdeaPayload
	if err := json.NewDecoder(r.Body).Decode(&updatePayload); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode idea: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.GetIdea(id)
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusNotFound)
		return
	}

	updatedIdea := result.Data

	if updatePayload.Title != nil {
		updatedIdea.Title = *updatePayload.Title
	}
	if updatePayload.Description != nil {
		updatedIdea.Description = *updatePayload.Description
	}
	if updatePayload.TechStack != nil {
		updatedIdea.TechStack = *updatePayload.TechStack
	}
	if updatePayload.Tags != nil {
		// tagsJSON, err := json.Marshal(updatePayload.Tags)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("failed to marshal tags: %v", err), http.StatusBadRequest)
		// 	return
		// }
		// updatedIdea.Tags = tagsJSON
		updatedIdea.Tags = *updatePayload.Tags
	}

	if updatePayload.Status != nil {
		updatedIdea.Status = *updatePayload.Status
	}

	if updatePayload.RequestedBy != nil {
		updatedIdea.RequestedBy = *updatePayload.RequestedBy
	}

	updatedIdea.UpdatedAt = time.Now()

	updateResult := h.service.UpdateIdea(id, updatedIdea)
	if updateResult.Err != nil {
		http.Error(w, updateResult.Err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": updateResult.Data})
}

// DeleteIdea godoc
// @Summary Delete an idea
// @Description Removes an idea from the system
// @Tags Ideas
// @Produce json
// @Param id path string true "Idea ID"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} error "Invalid ID format"
// @Failure 404 {object} error "Idea not found"
// @Failure 500 {object} error "Server error"
// @Router /idea/{id} [delete]
func (h *IdeaHandler) DeleteIdea(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	result := h.service.DeleteIdea(id)
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": result.Data})
}
