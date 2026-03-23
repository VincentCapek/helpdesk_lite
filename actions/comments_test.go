package actions

import (
	"helpdesk_lite/models"
	"net/http"
)

func (as *ActionSuite) Test_CommentsCreate_OK() {
	user := createActionUser(as, "comments-create-ok@example.com", "user")
	ticket := createActionTicket(as, user.ID)

	count, err := as.DB.Count("comments")
	as.NoError(err)
	as.Equal(0, count)

	comment := &models.Comment{
		Body: "I can reproduce this bug on Chrome as well.",
		Internal: false,
		UserID: user.ID,
	}

	res := as.HTML("/tickets/%s/comments", ticket.ID).Post(comment)

	as.Equal(http.StatusFound, res.Code)

	count, err = as.DB.Count("comments")
	as.NoError(err)
	as.Equal(1, count)

	created := &models.Comment{}
	err = as.DB.Where("body = ?", "I can reproduce this bug on Chrome as well.").First(created)
	as.NoError(err)
	as.Equal(user.ID, created.UserID)
	as.Equal(ticket.ID, created.TicketID)
	as.False(created.Internal)
}

func (as *ActionSuite) Test_CommentsCreate_ValidationErrors() {
	user := createActionUser(as, "comments-create-validation-errors@example.com", "user")
	ticket := createActionTicket(as, user.ID)

	count, err := as.DB.Count("comments")
	as.NoError(err)
	as.Equal(0, count)

	comment := &models.Comment{
		Body: "",
		Internal: false,
	}

	res := as.HTML("/tickets/%s/comments", ticket.ID).Post(comment)
	
	as.Equal(http.StatusUnprocessableEntity, res.Code)

	count, err = as.DB.Count("comments")
	as.NoError(err)
	as.Equal(0, count)

	as.Contains(res.Body.String(), "Please fix the following errors")
	as.Contains(res.Body.String(), ticket.Title)
}