package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/polarisbase/polarisbase-authn/internal/api/common"
	v1 "github.com/polarisbase/polarisbase-authn/internal/api/v1"
	"github.com/polarisbase/polarisbase-authn/internal/info"
	"github.com/polarisbase/polarisbase-authn/internal/user"
)

type Api struct {
	dep *common.Dependencies
	v1  *v1.Api
}

func New(fiberRouter fiber.Router, infoActionsProvider *info.ActionsProvider, userActionsProvider *user.ActionsProvider) *Api {

	// Create the authentication API
	a := &Api{}

	// add cors middleware
	fiberRouter.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Create the shared dependencies
	a.dep = common.NewDependencies(fiberRouter, infoActionsProvider, userActionsProvider)

	// Create the v1 API
	a.v1 = v1.New(a.dep)

	// Return the authentication API
	return a
}
