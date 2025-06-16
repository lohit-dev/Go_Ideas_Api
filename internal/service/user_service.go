package service

import (
	"errors"
	"fmt"
	"test_project/test/internal/model"
	"test_project/test/internal/storage"
	utils "test_project/test/pkg"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store storage.UserStorage
}

func NewUserService(store storage.UserStorage) *UserService {
	return &UserService{store}
}

func (s *UserService) CreateUser(req model.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	return s.store.CreateUser(user)
}

func (s *UserService) ValidateCredentials(username, password string) (bool, error) {
	user, err := s.store.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}

func (s *UserService) GetUserByUsername(username string) utils.Result[model.User] {
	user, err := s.store.GetUserByUsername(username)
	if err != nil {
		return utils.Result[model.User]{Err: err}
	}

	return utils.Result[model.User]{Data: user}
}

func (s *UserService) GetAllUsers() utils.Result[[]model.User] {
	return s.store.GetAllUsers()
}

func (s *UserService) DeleteUser(username, password string) (bool, error) {
	user, err := s.store.GetUserByUsername(username)
	if err != nil {
		return false, fmt.Errorf("failed to fetch user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, fmt.Errorf("invalid credentials")
	}

	_, err = s.store.DeleteUser(username)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	return true, nil
}
