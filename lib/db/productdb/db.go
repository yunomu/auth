package productdb

import (
	"context"
	"errors"
)

type Product struct {
	ClientId string
	AppCode  string
	FuncArn  string
}

var ErrNotFound = errors.New("not found")

type DB interface {
	Get(ctx context.Context, clientId string) (*Product, error)
	Scan(ctx context.Context, f func(*Product)) error
	Put(ctx context.Context, record *Product) error
}
