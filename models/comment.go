package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Comment is used by pop to map your comments database table to your go code.
type Comment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Body      string    `json:"body" db:"body"`
	Internal  bool      `json:"internal" db:"internal"`
	TicketID  uuid.UUID `json:"ticket_id" db:"ticket_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Ticket Ticket `belongs_to:"ticket"`
	User   User   `belongs_to:"user"`
}

// String is not required by pop and may be deleted
func (c Comment) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Comments is not required by pop and may be deleted
type Comments []Comment

// String is not required by pop and may be deleted
func (c Comments) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Comment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Body, Name: "Body"},
		&validators.FuncValidator{
			Field: string(c.TicketID.String()),
			Name: "TicketID",
			Message: "Ticket must be present",
			Fn: func () bool {
				return c.TicketID != uuid.Nil
			},
		},
		&validators.FuncValidator{
			Field: string(c.UserID.String()),
			Name: "UserID",
			Message: "User must be present",
			Fn: func () bool {
				return c.UserID != uuid.Nil
			},
		},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Comment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Comment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
