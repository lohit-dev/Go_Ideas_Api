package service

import (
	"fmt"
	"test_project/test/internal/model"
	"test_project/test/internal/storage"
	utils "test_project/test/pkg"

	"github.com/google/uuid"
)

type IdeaService struct {
	store storage.IdeaStorage
}

func NewIdeaService(store storage.IdeaStorage) *IdeaService {
	return &IdeaService{store}
}

func (s *IdeaService) CreateIdea(idea model.Idea) utils.Result[string] {
	for _, stack := range idea.TechStack {
		if !utils.IsValidTechStack(stack) {
			return utils.Result[string]{Err: fmt.Errorf("invalid tech stack: %v", stack)}
		}
	}

	return s.store.CreateIdea(idea)
}

func (s *IdeaService) GetAllIdeas() utils.Result[[]model.Idea] {
	return s.store.GetAllIdeas()
}

func (s *IdeaService) GetIdea(id uuid.UUID) utils.Result[model.Idea] {
	return s.store.GetIdea(id)
}
func (s *IdeaService) UpdateIdea(id uuid.UUID, idea model.Idea) utils.Result[string] {

	if !utils.IsValidRequestStatus(idea.Status) {
		return utils.Result[string]{Err: fmt.Errorf("invalid request status: %v", idea.Status)}
	}

	for _, stack := range idea.TechStack {
		if !utils.IsValidTechStack(stack) {
			return utils.Result[string]{Err: fmt.Errorf("invalid tech stack: %v", stack)}
		}
	}

	return s.store.UpdateIdea(id, idea)
}

func (s *IdeaService) DeleteIdea(id uuid.UUID) utils.Result[string] {
	return s.store.DeleteIdea(id)
}
