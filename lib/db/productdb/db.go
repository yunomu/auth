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
var ErrOptimisticLock = errors.New("lock error")

type DB interface {
	Get(ctx context.Context, clientId string) (*Product, int64, error)
	Scan(ctx context.Context, f func(*Product, int64)) error
	Put(ctx context.Context, record *Product) (int64, error)
	Update(ctx context.Context, record *Product, version int64) (int64, error)
	Delete(ctx context.Context, clientId string, version int64) error
}
