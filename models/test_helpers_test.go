package models_test

import (
	"fmt"
	"helpdesk_lite/models"

	"github.com/gofrs/uuid"
)

func createUser(ms *ModelSuite, email string, role string) *models.User {
	user := &models.User{
		Email: email,
		PasswordHash: "hashed-password",
		Role: role,
	}
	verrs, err := ms.DB.ValidateAndCreate(user)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(user.ID)

	return user
}

func createTicket(ms *ModelSuite, userID uuid.UUID) *models.Ticket {
	t := &models.Ticket{
		Title: "Login issue",
		Description: "I cannot sign in to my account",
		Status: "open",
		Priority: "medium",
		Category: "bug",
		UserID: userID,
	}
	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(t.ID)
	return t
}

func uniqueEmail(prefix string) string {
	return fmt.Sprintf("%s-%s@example.com", prefix, uuid.Must(uuid.NewV4()).String())
}