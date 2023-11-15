package basic_store

import (
	"context"
	"github.com/google/uuid"
	"github.com/polarisbase/polarisbase-authn/internal/user/model"
	persist "github.com/polarisbase/polarisbase-persist"
	"github.com/upper/db/v4"
	"strings"
)

type Store struct {
	bucket         persist.Bucket
	UserCollection db.Collection
}

func (s *Store) ListUsers(ctx context.Context, limit int, offset int) (users []model.User, err error, ok bool) {
	err = s.UserCollection.Find().All(&users)
	if err != nil {
		return users, err, false
	}
	return users, nil, true
}

func (s *Store) CreateUser(ctx context.Context, userIn model.User) (user model.User, err error, ok bool) {
	userIn.ID = uuid.New().String()
	err = s.UserCollection.InsertReturning(&userIn)
	if err != nil {
		return user, err, false
	}
	return userIn, nil, true
}

func (s *Store) CheckIfEmailIsAlreadyInUse(email string) (err error, ok bool) {
	var user model.User
	err = s.UserCollection.Find("email", email).One(&user)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, true
		}
		if strings.Contains(err.Error(), "no more rows in this result set") {
			return nil, true
		}
		return err, false
	}
	return nil, false
}

func (s *Store) LookupByID(getContext context.Context, id string) (user model.User, err error, ok bool) {
	err = s.UserCollection.Find("id", id).One(&user)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func (s *Store) LookupByEmail(getContext context.Context, email string) (user model.User, err error, ok bool) {
	err = s.UserCollection.Find("email", email).One(&user)
	if err != nil {
		return user, err, false
	}
	return user, nil, true
}

func New(bucket persist.Bucket) *Store {

	s := &Store{}

	s.bucket = bucket

	userCollection, err := bucket.Collection("user", model.User{})

	if err != nil {
		panic(err)
	}

	s.UserCollection = userCollection

	return s

}
