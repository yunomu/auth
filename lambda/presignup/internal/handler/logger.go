package handler

import "github.com/yunomu/auth/lib/db/userlist"

type Logger interface {
	Signup(user *userlist.User)
	Info(msg string, email string, appCode string)
}

type defaultLogger struct{}

var _ Logger = (*defaultLogger)(nil)

func (l *defaultLogger) Signup(*userlist.User)       {}
func (l *defaultLogger) Info(string, string, string) {}
