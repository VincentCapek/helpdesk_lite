package models_test

import (
	"helpdesk_lite/models"

	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Ticket_Create_OK() {
	user := createUser(ms, uniqueEmail("ticket-ok"), "user")
	
	count, err := ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Ticket{
		Title: "Checkout page broken",
		Description: "The payment form crashes on submit",
		Status: "open",
		Priority: "high",
		Category: "bug",
		UserID: user.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(t.ID)

	count, err = ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Ticket_Create_ValidationErrors() {
	user := createUser(ms, uniqueEmail("ticket-invalid"), "user")

	count, err := ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Ticket{
		Title:       "",
		Description: "",
		Status: "pending",
		Priority: "urgent",
		Category: "security",
		UserID: user.ID,
	}
	
	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Ticket_Create_RequiresUserID() {
	count, err := ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Ticket{
		Title: "Cannot update profile",
		Description: "Saving profile changes fails",
		Status: "open",
		Priority: "medium",
		Category: "bug",
		UserID: uuid.Nil,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tickets")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Ticket_Update_OK() {
	user := createUser(ms, uniqueEmail("ticket-update-ok"), "user")
	t := createTicket(ms, user.ID)

	t.Status = "in_progress"
	t.Priority = "high"

	verrs, err := ms.DB.ValidateAndUpdate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	reloaded := &models.Ticket{}
	err = ms.DB.Find(reloaded, t.ID)
	ms.NoError(err)
	ms.Equal("in_progress", reloaded.Status)
	ms.Equal("high", reloaded.Priority)
}

func (ms *ModelSuite) Test_Ticket_Update_ValidationErrors() {
	user := createUser(ms, uniqueEmail("ticket-update-invalid"), "user")
	t := createTicket(ms, user.ID)

	t.Status = "archived"

	verrs, err := ms.DB.ValidateAndUpdate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	reloaded := &models.Ticket{}
	err = ms.DB.Find(reloaded, t.ID)
	ms.NoError(err)
	ms.Equal("open", reloaded.Status)
}