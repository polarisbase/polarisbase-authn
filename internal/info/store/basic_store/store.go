package basic_store

import (
	"context"
	"github.com/google/uuid"
	"github.com/polarisbase/polarisbase-authn/internal/info/model"
	"github.com/polarisbase/polarisbase-persist/document"
)

type Store struct {
	persistenceStore document.Store
}

func (s *Store) List(ctx context.Context, limit int, offset int) (infos []model.Info, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().Model(&infos).Limit(limit).Offset(offset).Scan(ctx)
	if err != nil {
		return infos, err, false
	}
	return infos, nil, true
}

func (s *Store) CreateInfo(ctx context.Context, infoIn model.Info) (info model.Info, err error, ok bool) {
	infoIn.ID = uuid.New().String()
	b := s.persistenceStore.GetBun()
	_, err = b.NewInsert().Model(&infoIn).Exec(ctx)
	if err != nil {
		return info, err, false
	}
	return infoIn, nil, true
}

func New(persistenceStore document.Store) *Store {

	s := &Store{}

	s.persistenceStore = persistenceStore

	s.persistenceStore.MigrateUsing(
		(*model.Info)(nil),
	)

	return s

}
