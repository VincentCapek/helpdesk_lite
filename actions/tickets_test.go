package actions

import (
	"helpdesk_lite/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func createActionUser(as *ActionSuite, email string, role string) *models.User {
	user := &models.User{
		Email: email,
		PasswordHash: "hashed-password",
		Role: role,
	}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.NotZero(user.ID)
	return user
}

func createActionTicket(as *ActionSuite, userID uuid.UUID) *models.Ticket {
	ticket := &models.Ticket{
		Title: "Login issue",
		Description: "I cannot sign in to my account",
		Status: "open",
		Priority: "medium",
		Category: "bug",
		UserID: userID,
	}
	verrs, err := as.DB.ValidateAndCreate(ticket)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.NotZero(ticket.ID)
	return ticket
}

func (as *ActionSuite) Test_TicketsIndex_OK() {
	user := createActionUser(as, "tickets-index@example.com", "user")
	ticket := createActionTicket(as, user.ID)
	
	res := as.HTML("/tickets").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), ticket.Title)
}

func (as *ActionSuite) Test_TicketsNew() {
	res := as.HTML("/tickets/new").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Create a New Ticket")
}

func (as *ActionSuite) Test_TicketsShow() {
	user := createActionUser(as, "tickets-create@example.com", "user")
	ticket := createActionTicket(as, user.ID)
	
	res := as.HTML("/tickets/%s", ticket.ID).Get()
	
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), ticket.Title)
	as.Contains(res.Body.String(), ticket.Description)
}

func (as *ActionSuite) Test_TicketsCreate_OK() {
	user := createActionUser(as, "tickets-create-ok@example.com", "user")

	count, err := as.DB.Count("tickets")
	as.NoError(err)
	as.Equal(0, count)

	ticket := &models.Ticket{
		Title: "Checkout page broken",
		Description: "The payment form crashes on submit",
		Status: "open",
		Priority: "high",
		Category: "bug",
		UserID: user.ID,
	}
	
	res := as.HTML("/tickets").Post(ticket)

	as.Equal(http.StatusFound, res.Code)

	count, err = as.DB.Count("tickets")
	as.NoError(err)
	as.Equal(1, count)

	created := &models.Ticket{}
	err = as.DB.Where("title = ?", "Checkout page broken").First(created)
	as.NoError(err)
	as.Equal("high", created.Priority)
	as.Equal(user.ID, created.UserID)
}

func (as *ActionSuite) Test_TicketsCreate_ValidationErrors() {
	user := createActionUser(as, "tickets-create-validation-errors@example.com", "user")

	count, err := as.DB.Count("tickets")
	as.NoError(err)
	as.Equal(0, count)

	ticket := &models.Ticket{
		Title: "",
		Description: "",
		Status: "pending",
		Priority: "urgent",
		Category: "security",
		UserID: user.ID,
	}

	res := as.HTML("/tickets").Post(ticket)

	as.Equal(http.StatusUnprocessableEntity, res.Code)

	count, err = as.DB.Count("tickets")
	as.NoError(err)
	as.Equal(0, count)

	as.Contains(res.Body.String(), "Please fix the following errors")
}