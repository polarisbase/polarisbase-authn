package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-authn/internal/info"
	"github.com/polarisbase/polarisbase-authn/internal/user"
)

type Dependencies struct {
	FiberRouter fiber.Router
	InfoActions *info.ActionsProvider
	UserActions *user.ActionsProvider
}

func NewDependencies(fiberRouter fiber.Router, infoActionsProvider *info.ActionsProvider, userActionsProvider *user.ActionsProvider) *Dependencies {

	d := &Dependencies{}

	d.FiberRouter = fiberRouter

	d.InfoActions = infoActionsProvider

	d.UserActions = userActionsProvider

	return d

}
