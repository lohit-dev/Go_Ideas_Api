package model

import (
	"time"

	"github.com/google/uuid"
)

// RequestStatus represents the status of an idea
// @Description Status of the idea request
type RequestStatus string

// TechStack represents the technology stack
// @Description Technology stack for the idea
type TechStack string

const (
	Requested  RequestStatus = "requested"
	Reviewing  RequestStatus = "reviewing"
	Planned    RequestStatus = "planned"
	InProgress RequestStatus = "in-progress"
	Published  RequestStatus = "published"
	Rejected   RequestStatus = "rejected"
)

const (
	Rust      TechStack = "Rust"
	Go        TechStack = "Go"
	Next      TechStack = "Next"
	React     TechStack = "React"
	Axum      TechStack = "Axum"
	Postgres  TechStack = "Postgres"
	MySQL     TechStack = "MySQL"
	Docker    TechStack = "Docker"
	ActixWeb  TechStack = "ActixWeb"
	ChiRouter TechStack = "ChiRouter"
	Node      TechStack = "Node"
)

type Idea struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	TechStack   []TechStack   `json:"techStack"`
	Tags        []string      `json:"tags"`
	Status      RequestStatus `json:"status"`
	// Votes       int           `json:"votes"`       // Number of user votes/interest
	// RequestedBy     string        `json:"requestedBy"`     // Username or email of requester
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateIdeaPayload struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	TechStack   []TechStack   `json:"techStack"`
	Tags        []string      `json:"tags"`
	Status      RequestStatus `json:"status"`
}

type UpdateIdeaPayload struct {
	Title       *string        `json:"title,omitempty"` // Pointer so we can check for null values
	Description *string        `json:"description,omitempty"`
	TechStack   *[]TechStack   `json:"techStack,omitempty"`
	Tags        *[]string      `json:"tags,omitempty"`
	Status      *RequestStatus `json:"status,omitempty"`
}
