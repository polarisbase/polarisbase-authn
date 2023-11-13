package basic_store

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/polarisbase/polarisbase-authn/internal/user/model"
	"github.com/polarisbase/polarisbase-persist/document"
)

type Store struct {
	persistenceStore document.Store
	userTable        string
}

func (s *Store) ListUsers(ctx context.Context, limit int, offset int) (infos []model.User, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().
		Model(&infos).
		ModelTableExpr(fmt.Sprintf("%s AS user", s.userTable)).
		Limit(limit).
		Offset(offset).Scan(ctx)
	if err != nil {
		return infos, err, false
	}
	return infos, nil, true
}

func (s *Store) CreateUser(ctx context.Context, userIn model.User) (user model.User, err error, ok bool) {
	userIn.ID = uuid.New().String()
	b := s.persistenceStore.GetBun()
	_, err = b.NewInsert().
		Model(&userIn).
		ModelTableExpr(fmt.Sprintf("%s AS user", s.userTable)).
		Exec(ctx)
	if err != nil {
		return user, err, false
	}
	return userIn, nil, true
}

func (s *Store) CheckIfEmailIsAlreadyInUse(email string) (err error, ok bool) {
	b := s.persistenceStore.GetBun()
	var user model.User
	err = b.NewSelect().
		Model(&user).
		ModelTableExpr(fmt.Sprintf("%s AS user", s.userTable)).
		Where("email = ?", email).Scan(context.Background())
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
	err = b.NewSelect().
		Model(&user).
		ModelTableExpr(fmt.Sprintf("%s AS user", s.userTable)).
		Where("id = ?", id).Scan(getContext)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func (s *Store) LookupByEmail(getContext context.Context, email string) (user model.User, err error, ok bool) {
	b := s.persistenceStore.GetBun()
	err = b.NewSelect().
		Model(&user).
		ModelTableExpr(fmt.Sprintf("%s AS user", s.userTable)).
		Where("email = ?", email).Scan(getContext)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func New(namespace string, persistenceStore document.Store) *Store {

	s := &Store{}

	s.persistenceStore = persistenceStore

	tableName, err := s.persistenceStore.Migrate(
		namespace,
		(*model.User)(nil),
	)

	if err != nil {
		panic(err)
	}

	s.userTable = tableName

	return s

}
