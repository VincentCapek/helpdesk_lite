package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// AdminDashboard default implementation.
func AdminDashboard(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("admin/dashboard.html"))
}

