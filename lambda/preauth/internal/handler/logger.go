package handler

import (
	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
)

type Logger interface {
	Error(err error, msg string, req *Request)
	Info(msg string, req *Request, product *productdb.Product, user *userlist.User)
}

type defaultLogger struct{}

func (*defaultLogger) Error(error, string, *Request) {}

func (*defaultLogger) Info(string, *Request, *productdb.Product, *userlist.User) {
}
