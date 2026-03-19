package actions

import (
	"fmt"
	"helpdesk_lite/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// CommentsCreate default implementation.
func CommentsCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	ticket := &models.Ticket{}
	if err := tx.Eager("User", "Comments", "Comments.User").Find(ticket, c.Param("ticket_id")); err != nil {
		return errors.WithStack(err)
	}

	comment := &models.Comment{}
	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}

	comment.TicketID = ticket.ID

	if cu := currentUser(c); cu != nil {
		comment.UserID = cu.ID
	}

	verrs, err := comment.Validate(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("ticket", ticket)
		c.Set("comment", comment)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("tickets/show.plush.html"))
	}

	c.Flash().Add("success", "Comment created successfully.")
	return c.Redirect(http.StatusFound, fmt.Sprintf("/tickets/%s", ticket.ID))
}

