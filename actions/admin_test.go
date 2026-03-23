package actions

import (
	"helpdesk_lite/models"
	"net/http"
)

func (as *ActionSuite) Test_Admin_Dashboard() {
	user := createActionUser(as, "admin-dashboard@example.com", "admin")

	t1 := &models.Ticket{
		Title: "Open ticket",
		Description: "First open issue",
		Status: "open",
		Priority: "high",
		Category: "bug",
		UserID: user.ID,
	}

	t2 := &models.Ticket{
		Title: "Resolved ticket",
		Description: "Second resolved issue",
		Status: "resolved",
		Priority: "low",
		Category: "other",
		UserID: user.ID,
	}

	verrs, err := as.DB.ValidateAndCreate(t1)
	as.NoError(err)
	as.False(verrs.HasAny())

	verrs, err = as.DB.ValidateAndCreate(t2)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/admin").Get()

	as.Equal(http.StatusOK, res.Code, res.Body.String())
	as.Contains(res.Body.String(), "Admin Dashboard")
	as.Contains(res.Body.String(), "Open ticket")
	as.Contains(res.Body.String(), "Total Tickets")
	as.Contains(res.Body.String(), "Open Tickets")
	as.Contains(res.Body.String(), "High Priority")
	as.Contains(res.Body.String(), "Unassigned")
}

