package storage

import (
	"test_project/test/internal/model"
	utils "test_project/test/pkg"

	"github.com/google/uuid"
)

type IdeaStorage interface {
	GetAllIdeas() utils.Result[[]model.Idea]
	GetIdea(id uuid.UUID) utils.Result[model.Idea]
	CreateIdea(idea model.Idea) utils.Result[string]
	UpdateIdea(id uuid.UUID, idea model.Idea) utils.Result[string]
	DeleteIdea(id uuid.UUID) utils.Result[string]
}

type UserStorage interface {
	CreateUser(user model.User) error
	GetUserByUsername(username string) (model.User, error)
}
