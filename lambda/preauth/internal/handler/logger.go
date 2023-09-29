package handler

import (
	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
)

type Logger interface {
	Error(error, string)
	Info(string, *Request, *productdb.Product, *userlist.User)
}

type defaultLogger struct{}

func (*defaultLogger) Error(_ error, _ string) {}

func (*defaultLogger) Info(string, *Request, *productdb.Product, *userlist.User) {
}
