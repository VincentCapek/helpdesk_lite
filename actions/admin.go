package actions

import (
	"helpdesk_lite/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// AdminDashboard default implementation.
func AdminDashboard(c buffalo.Context) error {
    tx := c.Value("tx").(*pop.Connection)

	allTickets := models.Tickets{}
	if err := tx.Eager("User").All(&allTickets); err != nil {
		return errors.WithStack(err)
	}

	openTickets := models.Tickets{}
	if err := tx.Where("status = ?", "open").All(&openTickets); err != nil {
		return errors.WithStack(err)
	}

	highPriorityTickets := models.Tickets{}
	if err := tx.Where("priority = ?", "high").All(&highPriorityTickets); err != nil {
		return errors.WithStack(err)
	}

	unassignedTickets := models.Tickets{}
	if err := tx.Where("agent_id IS NULL OR agent_id = ''").All(&unassignedTickets); err != nil {
		return errors.WithStack(err)
	}

	c.Set("allTickets", allTickets)
	c.Set("openTickets", openTickets)
	c.Set("highPriorityTickets", highPriorityTickets)
	c.Set("unassignedTickets", unassignedTickets)

	c.Set("totalTickets", len(allTickets))

	return c.Render(http.StatusOK, r.HTML("admin/dashboard.plush.html"))
}

