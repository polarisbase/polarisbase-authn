package store

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user/model"
)

type UserStore interface {
	ListUsers(ctx context.Context, limit int, offset int) (users []model.User, err error, ok bool)
	CreateUser(ctx context.Context, userIn model.User) (user model.User, err error, ok bool)
	CheckIfEmailIsAlreadyInUse(email string) (err error, ok bool)
	LookupByID(getContext context.Context, id string) (user model.User, err error, ok bool)
	LookupByEmail(getContext context.Context, email string) (user model.User, err error, ok bool)
}
