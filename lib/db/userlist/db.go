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

type DB interface {
	Get(ctx context.Context, name string) (*User, error)
	Scan(ctx context.Context, f func(*User)) error
	Put(ctx context.Context, user *User) error
}
