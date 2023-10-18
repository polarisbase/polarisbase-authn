package v1

import (
	"fmt"
	"github.com/polarisbase/polarisbase-authn/internal/api/common"
)

type Api struct {
	prefix string
	dep    *common.Dependencies
}

func New(dependencies *common.Dependencies) *Api {

	a := &Api{}

	a.prefix = "v1"

	a.dep = dependencies

	// Login (POST)
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/login", a.prefix),
		a.Login,
	)

	// Login (GET)
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/login", a.prefix),
		a.LoginPage,
	)

	// Logout (POST)
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/logout", a.prefix),
		a.Logout,
	)

	// Logout (GET)
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/logout", a.prefix),
		a.Logout,
	)

	// User Create (POST)
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/user", a.prefix),
		a.CreateUser,
	)

	// Create an info item
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/_/create-info", a.prefix),
		a.CreateInfo,
	)

	// List the info items
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/_/list-info", a.prefix),
		a.ListInfo,
	)

	// Login Test
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/_/login-test", a.prefix),
		a.LoginTest,
	)

	return a

}
