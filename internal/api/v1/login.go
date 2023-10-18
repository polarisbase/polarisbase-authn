package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-authn/internal/user/contracts"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Api) Login(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// get request
	request := contracts.LoginRequest{}
	if err := c.BodyParser(&request); err != nil {
		c.Status(400)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	// validate login request
	if err := request.Validate(); err != nil {
		c.Status(400)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	// login
	response, err := a.dep.UserActions.V1.Login(ticket, request)
	if err != nil {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})
	}

	// set ticket user id
	err = ticket.AuthenticateTicket(response.User.ID, response.User)
	if err != nil {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
		})
	}

	// save ticket
	if err := ticket.SaveToCookies(c,
		func() (signingKey string, keyId string) {
			return "super-secret-key", "test-key-id"
		},
	); err != nil {
		c.Status(401)
		return c.JSON(map[string]string{
			"error": "unauthorized",
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
