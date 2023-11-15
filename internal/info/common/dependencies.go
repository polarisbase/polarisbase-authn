package common

import (
	"github.com/polarisbase/polarisbase-authn/internal/info/store"
	"github.com/polarisbase/polarisbase-authn/internal/info/store/basic_store"
	"github.com/polarisbase/polarisbase-persist"
)

type Dependencies struct {
	bucket    persist.Bucket
	InfoStore store.InfoStore
}

func NewDependencies(bucket persist.Bucket) *Dependencies {

	d := &Dependencies{}

	d.bucket = bucket

	d.InfoStore = basic_store.New(d.bucket)

	return d

}
