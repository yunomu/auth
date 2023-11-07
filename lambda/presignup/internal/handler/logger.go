package handler

import "github.com/yunomu/auth/lib/db/userlist"

type Logger interface {
	Signup(user *userlist.User)
	Info(msg string, email string, appCode string)
	GetProductError(err error, clientId string, req *Request)
}

type defaultLogger struct{}

var _ Logger = (*defaultLogger)(nil)

func (l *defaultLogger) Signup(*userlist.User)                   {}
func (l *defaultLogger) Info(string, string, string)             {}
func (l *defaultLogger) GetProductError(error, string, *Request) {}
