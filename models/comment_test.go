package models_test

import (
	"helpdesk_lite/models"

	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Comment_Create_OK() {
	user := createUser(ms, uniqueEmail("comment-ok"), "user")
	ticket := createTicket(ms, user.ID)

	count, err := ms.DB.Count("comments")
	ms.NoError(err)
	ms.Equal(0, count)

	comment := &models.Comment{
		Body: "I can reproduce this bug on Chrome as well.",
		Internal: false,
		TicketID: ticket.ID,
		UserID: user.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(comment)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(comment.ID)

	count, err = ms.DB.Count("comments")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Comment_Create_ValidationErrors() {
	count, err := ms.DB.Count("comments")
	ms.NoError(err)
	ms.Equal(0, count)

	comment := &models.Comment{
		Body: "",
		Internal: false,
		TicketID: uuid.Nil,
		UserID: uuid.Nil,
	}

	verrs, err := ms.DB.ValidateAndCreate(comment)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	
	count, err = ms.DB.Count("comments")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Comment_Update_OK() {
	user := createUser(ms, uniqueEmail("comment-update-ok"), "user")
	ticket := createTicket(ms, user.ID)
	
	comment := &models.Comment{
		Body: "Initial public reply",
		Internal: false,
		TicketID: ticket.ID,
		UserID: user.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(comment)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	comment.Body = "Updated internal note"
	comment.Internal = true

	verrs, err = ms.DB.ValidateAndUpdate(comment)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	reloaded := &models.Comment{}
	err = ms.DB.Find(reloaded, comment.ID)
	ms.NoError(err)
	ms.Equal("Updated internal note", reloaded.Body)
	ms.True(reloaded.Internal)
}

func (ms *ModelSuite) Test_Comment_Update_ValidationErrors() {
	user := createUser(ms, uniqueEmail("comment-update-invalid"), "user")
	ticket := createTicket(ms, user.ID)

	comment := &models.Comment{
		Body: "Some useful context",
		Internal: false,
		TicketID: ticket.ID,
		UserID: user.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(comment)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	
	comment.Body = ""

	verrs, err = ms.DB.ValidateAndUpdate(comment)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	reloaded := &models.Comment{}
	err = ms.DB.Find(reloaded, comment.ID)
	ms.NoError(err)
	ms.Equal("Some useful context", reloaded.Body)
}