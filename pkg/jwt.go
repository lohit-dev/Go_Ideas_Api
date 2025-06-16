package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func ExtractUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return uuid.Nil, errors.New("no auth header")
	}

	tokenStr := strings.Split(authHeader, " ")
	if len(tokenStr) != 2 {
		return uuid.Nil, errors.New("invalid token format")
	}

	token, err := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnvOrDefault("JWT_SECRET", "default")), nil
	})
	if err != nil || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("could not parse claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, errors.New("user_id not found in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user_id format in token")
	}

	return userID, nil
}
