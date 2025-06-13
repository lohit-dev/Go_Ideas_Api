package handler

import (
	"encoding/json"
	"net/http"
	"test_project/test/internal/model"
	"test_project/test/internal/service"
	utils "test_project/test/pkg"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	userService *service.UserService
	validator   *validator.Validate
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.CreateUser(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, err := h.userService.ValidateCredentials(req.Username, req.Password)
	if err != nil || !valid {
		h.sendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := utils.GetEnvOrDefault("JWT_SECRET", "default")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.LoginResponse{Token: tokenString})
}

func (h *AuthHandler) sendError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *AuthHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleted, err := h.userService.DeleteUser(req.Username, req.Password)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if deleted {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
		return
	}

	h.sendError(w, "Failed to delete user", http.StatusInternalServerError)
}

func (h *AuthHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := h.userService.GetAllUsers()
	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Data)
}

func (h *AuthHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "missing username query parameter", http.StatusBadRequest)
		return
	}

	result := h.userService.GetUserByUsername(username)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Data)
}
