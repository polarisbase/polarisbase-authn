package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Api) Logout(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// logout
	err := ticket.RemoveFromCookies(c)
	if err != nil {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})
	}

	c.Status(200)
	return c.JSON(map[string]string{
		"message": "logged out",
	})

}
