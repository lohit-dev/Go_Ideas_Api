package storage

import (
	"errors"
	"fmt"
	"test_project/test/internal"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ps *PostgresStore) CreateUser(user internal.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}

	var existingUser internal.User
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

func (ps *PostgresStore) GetUserByUsername(username string) (internal.User, error) {
	var user internal.User

	if err := ps.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return internal.User{}, errors.New("user not found")
		}

		return internal.User{}, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil
}
