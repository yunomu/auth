package handler

import "github.com/yunomu/auth/lib/userlist"

type Logger interface {
	Signup(*userlist.User)
}

type defaultLogger struct{}

var _ Logger = (*defaultLogger)(nil)

func (l *defaultLogger) Signup(*userlist.User) {}
