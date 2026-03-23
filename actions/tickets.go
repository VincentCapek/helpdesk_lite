package actions

import (
	"fmt"
	"net/http"

	"helpdesk_lite/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// TicketsIndex default implementation.
func TicketsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	tickets := models.Tickets{}
	if err := tx.Eager("User").All(&tickets); err != nil {
		return errors.WithStack(err)
	}

	c.Set("tickets", tickets)
	return c.Render(http.StatusOK, r.HTML("tickets/index.plush.html"))
}

// TicketsShow default implementation.
func TicketsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	ticket := &models.Ticket{}
	if err := tx.Eager("User", "Comments", "Comments.User").Find(ticket, c.Param("ticket_id")); err != nil {
		return errors.WithStack(err)
	}

	comment := &models.Comment{}
	c.Set("ticket", ticket)
	c.Set("comment", comment)
	c.Set("commentErrors", nil)
	return c.Render(http.StatusOK, r.HTML("tickets/show.plush.html"))
}

// TicketsNew default implementation.
func TicketsNew(c buffalo.Context) error {
    ticket := &models.Ticket{
		Status: "open",
		Priority: "medium",
		Category: "other",
	}

	c.Set("ticket", ticket)
	return c.Render(http.StatusOK, r.HTML("tickets/new.plush.html"))
}

// TicketsCreate default implementation.
func TicketsCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	ticket := &models.Ticket{}
	if err := c.Bind(ticket); err != nil {
		return errors.WithStack(err)
	}

	if cu := currentUser(c); cu != nil {
		ticket.UserID = cu.ID
	}

	verrs, err := ticket.Validate(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("ticket", ticket)
		c.Set("ticketErrors", verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("tickets/new.plush.html"))
	}

	if err := tx.Create(ticket); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "Ticket created successfully.")
	return c.Redirect(http.StatusFound, fmt.Sprintf("/tickets/%s", ticket.ID))
}

// TicketsEdit default implementation.
func TicketsEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	ticket := &models.Ticket{}
	if err := tx.Find(ticket, c.Param("ticket_id")); err != nil {
		return errors.WithStack(err)
	}

	c.Set("ticket", ticket)
	return c.Render(http.StatusOK, r.HTML("tickets/edit.plush.html"))
}

// TicketsUpdate default implementation.
func TicketsUpdate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	ticket := &models.Ticket{}
	if err := tx.Find(ticket, c.Param("ticket_id")); err != nil {
		return errors.WithStack(err)
	}

	if err := c.Bind(ticket); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := ticket.ValidateUpdate(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("ticket", ticket)
		c.Set("errors", verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("tickets/edit.plush.html"))
	}

	c.Flash().Add("success", "Ticket updated successfully.")
	return c.Redirect(http.StatusFound, fmt.Sprintf("/tickets/%s", ticket.ID))
}

// TicketsDestroy default implementation.
func TicketsDestroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	ticket := &models.Ticket{}
	if err := tx.Find(ticket, c.Param("ticket_id")); err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Destroy(ticket); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "Ticket deleted successfully.")
	return c.Render(http.StatusOK, r.HTML("tickets/destroy.plush.html"))
}

