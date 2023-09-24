package handler

import "github.com/yunomu/auth/lib/db/userlist"

type Logger interface {
	Signup(*userlist.User)
}

type defaultLogger struct{}

var _ Logger = (*defaultLogger)(nil)

func (l *defaultLogger) Signup(*userlist.User) {}
