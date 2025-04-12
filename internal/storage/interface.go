package storage

import (
	"encoding/json"
	"fmt"
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

// Implementation for json store
func (js *JsonStore) GetAllIdeas() utils.Result[[]model.Idea] {
	return js.ReadFile()
}

func (js *JsonStore) GetIdea(id uuid.UUID) utils.Result[model.Idea] {
	result := js.ReadFile()
	if result.Err != nil {
		return utils.Result[model.Idea]{Err: result.Err}
	}

	for _, idea := range result.Data {
		if idea.ID == id {
			return utils.Result[model.Idea]{Data: idea}
		}
	}

	return utils.Result[model.Idea]{Err: fmt.Errorf(`idea with ID %s not found`, id)}
}

func (js *JsonStore) CreateIdea(idea model.Idea) utils.Result[string] {
	result := js.ReadFile()
	if result.Err != nil {
		return utils.Result[string]{Err: result.Err}
	}

	idea.ID = uuid.MustParse(utils.GenId())
	ideas := append(result.Data, idea)

	data, err := json.Marshal(ideas)
	if err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to marshal ideas: %v", err)}
	}

	writeResult := js.WriteJson(js.filepath, data)
	if writeResult.Err != nil {
		return writeResult
	}

	return utils.Result[string]{Data: "Idea created successfully"}
}

func (js *JsonStore) UpdateIdea(id uuid.UUID, updatedIdea model.Idea) utils.Result[string] {
	result := js.ReadFile()
	if result.Err != nil {
		return utils.Result[string]{Err: result.Err}
	}

	for i, idea := range result.Data {
		if idea.ID == id {
			result.Data[i] = updatedIdea

			// Marshal the updated list to JSON
			data, err := json.Marshal(result.Data)
			if err != nil {
				return utils.Result[string]{Err: fmt.Errorf("failed to marshal ideas: %v", err)}
			}

			// Writing back
			writeResult := js.WriteJson(js.filepath, data)
			if writeResult.Err != nil {
				return writeResult
			}

			return utils.Result[string]{Data: "Idea updated successfully"}
		}
	}

	return utils.Result[string]{Err: fmt.Errorf("idea with ID %s not found", id)}
}

func (js *JsonStore) DeleteIdea(id uuid.UUID) utils.Result[string] {
	result := js.ReadFile()
	if result.Err != nil {
		return utils.Result[string]{Err: result.Err}
	}

	var updatedIdeas []model.Idea
	var found bool

	for _, idea := range result.Data {
		if idea.ID != id {
			updatedIdeas = append(updatedIdeas, idea)
		} else {
			found = true
		}
	}

	if !found {
		return utils.Result[string]{Err: fmt.Errorf("idea with ID %s not found", id)}
	}

	// If no ideas remain, write an empty array
	if len(updatedIdeas) == 0 {
		updatedIdeas = []model.Idea{}
	}

	data, err := json.Marshal(updatedIdeas)
	if err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to marshal ideas: %v", err)}
	}

	writeResult := js.WriteJson(js.filepath, data)
	if writeResult.Err != nil {
		return writeResult
	}

	return utils.Result[string]{Data: "Idea deleted successfully"}
}
