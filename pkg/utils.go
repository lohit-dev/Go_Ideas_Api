package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"test_project/test/internal/model"

	"github.com/google/uuid"
)

func GenId() string {
	return uuid.NewString()
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func IsValidTechStack(stack model.TechStack) bool {
	switch stack {
	case model.Rust, model.Go, model.Next, model.React, model.Axum, model.Postgres,
		model.MySQL, model.Docker, model.ActixWeb, model.ChiRouter, model.Node:
		return true
	default:
		return false
	}
}

func IsValidRequestStatus(status model.RequestStatus) bool {
	switch status {
	case model.Requested, model.Reviewing, model.Planned, model.InProgress, model.Published, model.Rejected:
		return true
	default:
		return false
	}
}

func UnmarshalJson(data []byte, result any) error {
	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}
	return nil
}

type Result[T any] struct {
	Data T
	Err  error
}

// func NewResult[T any](data T, err error) Result[T] {
// 	return Result[T]{Data: data, Err: err}
// }

func (r *Result[T]) New(data T, err error) Result[T] {
	return Result[T]{Data: data, Err: err}
}

func NewResult[T any](data T, err error) Result[T] {
	return Result[T]{Data: data, Err: err}
}
