package authn

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polarisbase-authn/internal/api"
	"github.com/polarisbase/polarisbase-authn/internal/info"
	"github.com/polarisbase/polarisbase-authn/internal/user"
	persist "github.com/polarisbase/polarisbase-persist"
)

type Service struct {
	// fiberRouter is the router for the authentication API
	fiberRouter fiber.Router
	// authApiPrefix is the prefix for all authentication routes
	authApiPrefix string
	// store is the document store
	store persist.Store
	// Bucket is the document bucket
	Bucket persist.Bucket
	// infoActionsProvider is the provider for info actions
	infoActionsProvider *info.ActionsProvider
	// userActionsProvider is the provider for user actions
	userActionsProvider *user.ActionsProvider
	// api is the authentication API
	api *api.Api
}

func New(fiberRouter fiber.Router, namespace string, authApiPrefix string, store persist.Store) *Service {
	// Create the authentication service
	s := &Service{
		authApiPrefix: authApiPrefix,
	}
	// Create the document store
	s.store = store
	// Create the document bucket
	s.Bucket, _ = store.NewBucket(namespace)
	// Create a sub-router for the authentication API
	s.fiberRouter = fiberRouter.Group(s.authApiPrefix)
	// Create the info actions provider
	s.infoActionsProvider = info.NewActionsProvider(s.Bucket)
	// Create the user actions provider
	s.userActionsProvider = user.NewActionsProvider(s.Bucket)
	// Create the authentication API
	s.api = api.New(s.fiberRouter, s.infoActionsProvider, s.userActionsProvider)
	// Return the authentication service
	return s
}
