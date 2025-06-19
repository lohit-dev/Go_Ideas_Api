package storage

import (
	"errors"
	"fmt"
	"test_project/test/internal/model"
	utils "test_project/test/pkg"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(dbstring string) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the db: %v", err)
	}

	// We must add the models here for creation of the tables
	if err := db.AutoMigrate(&model.Idea{}, &model.User{}, &model.Vote{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return &PostgresStore{db: db}, nil
}

func (ps *PostgresStore) GetAllIdeas() utils.Result[[]model.Idea] {
	var ideas []model.Idea
	// Preload Votes so the slice is filled
	if err := ps.db.Preload("Votes").Find(&ideas).Error; err != nil {
		return utils.Result[[]model.Idea]{Err: fmt.Errorf("failed to get all ideas: %v", err)}
	}
	// Set Count for each idea
	for i := range ideas {
		ideas[i].Count = len(ideas[i].Votes)
	}
	return utils.Result[[]model.Idea]{Data: ideas}
}

func (ps *PostgresStore) GetIdea(id uuid.UUID) utils.Result[model.Idea] {
	var idea model.Idea
	if err := ps.db.Find(&idea, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Result[model.Idea]{Err: fmt.Errorf("idea with that id %s is not fount", id)}
		}
		return utils.Result[model.Idea]{Err: fmt.Errorf("failed to get idea: %v", err)}
	}

	return utils.Result[model.Idea]{Data: idea}
}

func (ps *PostgresStore) CreateIdea(idea model.Idea) utils.Result[string] {
	if idea.ID == uuid.Nil {
		idea.ID = uuid.New()
	}

	if idea.Status == "" {
		idea.Status = model.Requested
	}

	if idea.CreatedAt.IsZero() {
		idea.CreatedAt = time.Now()
	}

	if idea.UpdatedAt.IsZero() {
		idea.UpdatedAt = time.Now()
	}

	if err := ps.db.Create(&idea).Error; err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to create an idea: %v", err)}
	}

	return utils.Result[string]{Data: "idea created successfully"}
}

func (ps *PostgresStore) UpdateIdea(id uuid.UUID, updatedIdea model.Idea) utils.Result[string] {
	var existing model.Idea
	if err := ps.db.First(&existing, "id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Result[string]{Err: fmt.Errorf("idea with id %s not found", id)}
		}
		return utils.Result[string]{Err: fmt.Errorf("failed to get the idea: %v", err)}
	}

	updatedIdea.UpdatedAt = time.Now()
	updatedIdea.ID = id

	if err := ps.db.Model(&existing).Updates(updatedIdea).Error; err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to update the idea: %v", err)}
	}

	return utils.Result[string]{Data: "Idea updated successfully"}
}

func (ps *PostgresStore) DeleteIdea(id uuid.UUID) utils.Result[string] {
	var existing model.Idea
	if err := ps.db.First(&existing, "id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Result[string]{Err: fmt.Errorf("idea with id %s not found", id)}
		}
		return utils.Result[string]{Err: fmt.Errorf("failed to get the idea: %v", err)}
	}

	if err := ps.db.Delete(&existing).Error; err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to delete the idea: %v", err)}
	}

	return utils.Result[string]{Data: "Idea deleted successfully"}
}
