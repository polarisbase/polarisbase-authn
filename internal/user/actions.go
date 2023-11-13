package user

import (
	"github.com/polarisbase/polarisbase-authn/internal/user/common"
	v1 "github.com/polarisbase/polarisbase-authn/internal/user/v1"
	"github.com/polarisbase/polarisbase-persist"
)

type ActionsProvider struct {
	dep *common.Dependencies
	V1  *v1.Actions
}

func NewActionsProvider(namespace string, persist persist.Persist) *ActionsProvider {
	// Create the actions provider
	ap := &ActionsProvider{}
	// Create the shared dependencies
	ap.dep = common.NewDependencies(namespace, persist)
	// Create the v1 actions
	ap.V1 = v1.New(ap.dep)
	// Return the actions provider
	return ap

}
