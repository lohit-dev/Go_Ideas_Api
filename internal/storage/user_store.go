package storage

import (
	"errors"
	"fmt"
	"test_project/test/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ps *PostgresStore) CreateUser(user model.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}

	var existingUser model.User
	result := ps.db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == nil {
		return errors.New("username already exists")
	}

	result = ps.db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return errors.New("email already exists")
	}

	if err := ps.db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (ps *PostgresStore) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	if err := ps.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}

		return model.User{}, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil
}
