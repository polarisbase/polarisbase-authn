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

func NewActionsProvider(bucket persist.Bucket) *ActionsProvider {
	// Create the actions provider
	ap := &ActionsProvider{}
	// Create the shared dependencies
	ap.dep = common.NewDependencies(bucket)
	// Create the v1 actions
	ap.V1 = v1.New(ap.dep)
	// Return the actions provider
	return ap

}
