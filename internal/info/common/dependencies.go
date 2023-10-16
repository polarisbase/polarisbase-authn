package common

import (
	"github.com/polarisbase/polaris-sdk/v3/lib/persist"
	"github.com/polarisbase/polaris-sdk/v3/lib/persist/document"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info/store"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info/store/basic_store"
)

type Dependencies struct {
	persist   persist.Persist
	InfoStore store.InfoStore
}

func NewDependencies(persist persist.Persist) *Dependencies {

	d := &Dependencies{}

	d.persist = persist

	// try and cast the persist document.store
	if documentStore, ok := persist.(document.Store); ok {
		d.InfoStore = basic_store.New(documentStore)
		return d
	} else {
		panic("persist is not a document store, cannot create info store dependency for authn service")
	}

	return d

}
