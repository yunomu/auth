package userlist

import (
	"context"
	"errors"
)

type User struct {
	Name     string
	AppCodes []string
}

var ErrNoSuchUser = errors.New("no such user")
var ErrOptimisticLock = errors.New("lock error")

type DB interface {
	Get(ctx context.Context, name string) (*User, int64, error)
	Scan(ctx context.Context, f func(*User, int64)) error
	Put(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User, timestamp int64) (int64, error)
	Delete(ctx context.Context, name string, timestamp int64) error
}
