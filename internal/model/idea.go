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
	ID          uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string        `json:"title" gorm:"not null"`
	Description string        `json:"description" gorm:"type:text"`
	TechStack   []TechStack   `json:"techStack" gorm:"type:jsonb"`
	Tags        []string      `json:"tags" gorm:"type:jsonb"`
	Status      RequestStatus `json:"status" gorm:"type:varchar(20);default:'requested'"`
	Votes       int           `json:"votes" gorm:"default:0"`
	RequestedBy string        `json:"requestedBy" gorm:"type:varchar(100)"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `json:"updatedAt" gorm:"autoUpdateTime"`
}

type CreateIdeaPayload struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	TechStack   []TechStack   `json:"techStack"`
	Tags        []string      `json:"tags"`
	Status      RequestStatus `json:"status,omitempty"`
	RequestedBy string        `json:"requestedBy,omitempty"`
}

type UpdateIdeaPayload struct {
	Title       *string        `json:"title,omitempty"`
	Description *string        `json:"description,omitempty"`
	TechStack   *[]TechStack   `json:"techStack,omitempty"`
	Tags        *[]string      `json:"tags,omitempty"`
	Status      *RequestStatus `json:"status,omitempty"`
	Votes       *int           `json:"votes,omitempty"`
	RequestedBy *string        `json:"requestedBy,omitempty"`
}
