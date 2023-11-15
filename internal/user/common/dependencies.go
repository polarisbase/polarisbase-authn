package common

import (
	"github.com/polarisbase/polarisbase-authn/internal/user/store"
	"github.com/polarisbase/polarisbase-authn/internal/user/store/basic_store"
	"github.com/polarisbase/polarisbase-persist"
)

type Dependencies struct {
	bucket    persist.Bucket
	UserStore store.UserStore
}

func NewDependencies(bucket persist.Bucket) *Dependencies {

	d := &Dependencies{}

	d.bucket = bucket

	d.UserStore = basic_store.New(d.bucket)

	return d

}
