package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type TicketPriority string

const (
	PriorityLow TicketPriority = "low"
	PriorityMedium TicketPriority = "medium"
	PriorityHigh TicketPriority = "high"
)

// Ticket is used by pop to map your tickets database table to your go code.
type Ticket struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Priority    string    `json:"priority" db:"priority"`
	Category    string    `json:"category" db:"category"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	AgentID     uuid.UUID `json:"agent_id" db:"agent_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	User      User      `belongs_to:"user"`
	Comments  Comments  `has_many:"comments"`
}

// String is not required by pop and may be deleted
func (t Ticket) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tickets is not required by pop and may be deleted
type Tickets []Ticket

// String is not required by pop and may be deleted
func (t Tickets) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Ticket) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
		&validators.StringInclusion{
			Field: t.Status,
			Name: "Status",
			List: []string{"open", "in_progress", "resolved", "closed"},
			Message: "Status must be one of: open, in_progress, resolved, closed",
		},
		&validators.StringInclusion{
			Field: t.Priority,
			Name: "Priority",
			List: []string{string(PriorityLow), string(PriorityMedium), string(PriorityHigh)},
			Message: "Priority must be one of: low, medium, high",
		},
		&validators.StringInclusion{
			Field: t.Category,
			Name: "Category",
			List: []string{"bug", "billing", "feature_request", "other"},
			Message: "Category must be one of: bug, billing, feature_request, other",
		},
		&validators.FuncValidator{
			Field: string(t.UserID.String()),
			Name: "UserID",
			Message: "User not found",
			Fn: func () bool {
				return t.UserID != uuid.Nil
			},
		},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Ticket) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Ticket) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
