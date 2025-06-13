package storage

import (
	"errors"
	"fmt"
	"test_project/test/internal/model"
	utils "test_project/test/pkg"
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
	}

	if err := ps.db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (ps *PostgresStore) GetAllUsers() utils.Result[[]model.User] {
	var users []model.User
	if err := ps.db.First(&users).Error; err != nil {
		return utils.Result[[]model.User]{Err: fmt.Errorf("failed to get all ideas: %v", err)}
	}

	return utils.Result[[]model.User]{Data: users}
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

func (ps *PostgresStore) DeleteUser(username string) (model.User, error) {
	var user model.User

	// Check if user exists
	result := ps.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, result.Error
	}

	if err := ps.db.Delete(&user).Error; err != nil {
		return model.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	return user, nil
}
