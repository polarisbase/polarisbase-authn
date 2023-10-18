package common

import (
	"github.com/polarisbase/polarisbase-authn/internal/user/store"
	"github.com/polarisbase/polarisbase-authn/internal/user/store/basic_store"
	"github.com/polarisbase/polarisbase-persist"
	"github.com/polarisbase/polarisbase-persist/document"
)

type Dependencies struct {
	persist   persist.Persist
	UserStore store.UserStore
}

func NewDependencies(persist persist.Persist) *Dependencies {

	d := &Dependencies{}

	d.persist = persist

	// try and cast the persist document.store
	if documentStore, ok := persist.(document.Store); ok {
		d.UserStore = basic_store.New(documentStore)
		return d
	} else {
		panic("persist is not a document store, cannot create user store dependency for authn service")
	}

	return d

}
