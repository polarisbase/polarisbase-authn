package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Api) LoginTest(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// set ticket user id
	ticket.AuthenticateTicket("user-id", nil)

	// save ticket
	if err := ticket.SaveToCookies(c,
		func() (signingKey string, keyId string) {
			return "super-secret-key", "test-key-id"
		},
	); err != nil {
		c.Status(500)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	// Do if user is authenticated
	return ticket.DoIfAuthenticated(func() error {
		return c.JSON(map[string]string{
			"message": "authenticated",
		})
	}, func() error {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})
	})

}
