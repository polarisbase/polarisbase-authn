package common

import (
	"github.com/polarisbase/polarisbase-authn/internal/info/store"
	"github.com/polarisbase/polarisbase-authn/internal/info/store/basic_store"
	"github.com/polarisbase/polarisbase-persist"
	"github.com/polarisbase/polarisbase-persist/document"
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
