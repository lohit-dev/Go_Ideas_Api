package model

import (
	"encoding/json"
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
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string          `json:"title" gorm:"not null"`
	Description string          `json:"description" gorm:"type:text"`
	TechStack   json.RawMessage `json:"techStack" gorm:"type:jsonb"`
	Tags        json.RawMessage `json:"tags" gorm:"type:jsonb"`
	Status      RequestStatus   `json:"status" gorm:"type:varchar(20);default:'requested'"`
	Votes       []Vote          `json:"votes,omitempty" gorm:"not null;foreignKey:IdeaID;default:0"`
	VoteCount   int             `json:"voteCount" gorm:"-"`
	RequestedBy string          `json:"requestedBy" gorm:"type:varchar(100)"`
	CreatedAt   time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
}

type CreateIdeaPayload struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	TechStack   json.RawMessage `json:"techStack"`
	Tags        json.RawMessage `json:"tags"`
	Status      RequestStatus   `json:"status,omitempty"`
	RequestedBy string          `json:"requestedBy,omitempty"`
}

type UpdateIdeaPayload struct {
	Title       *string          `json:"title,omitempty"`
	Description *string          `json:"description,omitempty"`
	TechStack   *json.RawMessage `json:"techStack,omitempty"`
	Tags        *json.RawMessage `json:"tags,omitempty"`
	Status      *RequestStatus   `json:"status,omitempty"`
	RequestedBy *string          `json:"requestedBy,omitempty"`
}

// *********************
//    voting related
// *********************

type Vote struct {
	ID        string    `json:"id"`
	IdeaID    uuid.UUID `json:"-"`
	UserID    uuid.UUID `json:"-"`
	User      User      `json:"user" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
}
