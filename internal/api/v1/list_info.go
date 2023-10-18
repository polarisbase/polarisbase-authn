package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-authn/internal/info/contracts"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Api) ListInfo(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// Do if user is authenticated
	return ticket.DoIfAuthenticated(func() error {

		res, err := a.dep.InfoActions.V1.List(ticket, contracts.ListRequest{
			Limit: 10,
		})

		if err != nil {
			c.Status(500)
			return c.JSON(map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(res)
	}, func() error {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})
	})

}
