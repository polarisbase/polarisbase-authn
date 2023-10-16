package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info/contracts"
	"github.com/polarisbase/polaris-sdk/v3/shared"
)

func (a *Api) CreateInfo(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// Do if user is authenticated
	return ticket.DoIfAuthenticated(func() error {

		res, err := a.dep.InfoActions.V1.Create(ticket, contracts.CreateRequest{})
		if err != nil {
			return err
		}

		c.Status(200)
		return c.JSON(map[string]string{
			"id": res.ID,
		})

	}, func() error {

		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})

	})

}
