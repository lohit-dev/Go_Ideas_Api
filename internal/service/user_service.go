package service

import (
	"errors"
	"test_project/test/internal"
	"test_project/test/internal/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store storage.UserStorage
}

func NewUserService(store storage.UserStorage) *UserService {
	return &UserService{store}
}

func (s *UserService) CreateUser(req internal.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := internal.User{
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
