package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-authn/internal/user/contracts"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Api) CreateUser(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// get request
	request := contracts.CreateRequest{}
	if err := c.BodyParser(&request); err != nil {
		c.Status(400)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	// validate create user request
	if err := request.Validate(); err != nil {
		c.Status(400)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	// create user
	response, err := a.dep.UserActions.V1.Create(ticket, request)
	if err != nil {
		c.Status(500)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(response)
}
