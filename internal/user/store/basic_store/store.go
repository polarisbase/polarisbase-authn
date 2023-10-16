package basic_store

import (
	"context"
	"github.com/google/uuid"
	"github.com/polarisbase/polaris-sdk/v3/lib/persist/document"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user/model"
)

type Store struct {
	persistenceStore document.Store
}

func (s *Store) ListUsers(ctx context.Context, limit int, offset int) (infos []model.User, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().Model(&infos).Limit(limit).Offset(offset).Scan(ctx)
	if err != nil {
		return infos, err, false
	}
	return infos, nil, true
}

func (s *Store) CreateUser(ctx context.Context, userIn model.User) (user model.User, err error, ok bool) {
	userIn.ID = uuid.New().String()
	b := s.persistenceStore.GetBun()
	_, err = b.NewInsert().Model(&userIn).Exec(ctx)
	if err != nil {
		return user, err, false
	}
	return userIn, nil, true
}

func (s *Store) CheckIfEmailIsAlreadyInUse(email string) (err error, ok bool) {
	b := s.persistenceStore.GetBun()
	var user model.User
	err = b.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, true
		}
		return err, false
	}
	return nil, false
}

func (s *Store) LookupByID(getContext context.Context, id string) (user model.User, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().Model(&user).Where("id = ?", id).Scan(getContext)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func (s *Store) LookupByEmail(getContext context.Context, email string) (user model.User, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().Model(&user).Where("email = ?", email).Scan(getContext)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func New(persistenceStore document.Store) *Store {

	s := &Store{}

	s.persistenceStore = persistenceStore

	s.persistenceStore.MigrateUsing(
		(*model.User)(nil),
	)

	return s

}
