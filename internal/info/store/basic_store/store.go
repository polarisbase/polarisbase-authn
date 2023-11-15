package basic_store

import (
	"context"
	"github.com/google/uuid"
	"github.com/polarisbase/polarisbase-authn/internal/info/model"
	persist "github.com/polarisbase/polarisbase-persist"
	"github.com/upper/db/v4"
)

type Store struct {
	bucket         persist.Bucket
	InfoCollection db.Collection
}

func (s *Store) List(ctx context.Context, limit int, offset int) (infos []model.Info, err error, ok bool) {
	err = s.InfoCollection.Find().All(&infos)
	if err != nil {
		return infos, err, false
	}
	return infos, nil, true
}

func (s *Store) CreateInfo(ctx context.Context, infoIn model.Info) (info model.Info, err error, ok bool) {
	infoIn.ID = uuid.New().String()
	err = s.InfoCollection.InsertReturning(&infoIn)
	if err != nil {
		return info, err, false
	}
	return infoIn, nil, true
}

func New(bucket persist.Bucket) *Store {

	s := &Store{}

	s.bucket = bucket

	infoCollection, err := bucket.Collection("info", &model.Info{})
	if err != nil {
		panic(err)
	}

	s.InfoCollection = infoCollection

	return s

}
