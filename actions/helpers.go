package actions

import (
	"helpdesk_lite/models"

	"github.com/gobuffalo/buffalo"
)

func currentUser(c buffalo.Context) *models.User {
	if v := c.Value("current_user"); v != nil {
		if u, ok := v.(*models.User); ok {
			return u
		}
	}
	return nil
}