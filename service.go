package authn

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/lib/persist/document"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/api"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user"
)

type Service struct {
	// fiberRouter is the router for the authentication API
	fiberRouter fiber.Router
	// authApiPrefix is the prefix for all authentication routes
	authApiPrefix string
	// documentStore is the document store
	documentStore document.Store
	// infoActionsProvider is the provider for info actions
	infoActionsProvider *info.ActionsProvider
	// userActionsProvider is the provider for user actions
	userActionsProvider *user.ActionsProvider
	// api is the authentication API
	api *api.Api
}

func New(fiberRouter fiber.Router, authApiPrefix string, documentStore document.Store) *Service {
	// Create the authentication service
	s := &Service{
		authApiPrefix: authApiPrefix,
	}
	// Create the document store
	s.documentStore = documentStore
	// Create a sub-router for the authentication API
	s.fiberRouter = fiberRouter.Group(s.authApiPrefix)
	// Create the info actions provider
	s.infoActionsProvider = info.NewActionsProvider(s.documentStore)
	// Create the user actions provider
	s.userActionsProvider = user.NewActionsProvider(s.documentStore)
	// Create the authentication API
	s.api = api.New(s.fiberRouter, s.infoActionsProvider, s.userActionsProvider)
	// Return the authentication service
	return s
}
