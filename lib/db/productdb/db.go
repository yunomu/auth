package productdb

import (
	"context"
	"errors"
)

type Record struct {
	ClientId string `dynamodbav:"ClientId"`
	AppCode  string `dynamodbav:"AppCode,omitempty"`
	Created  int64  `dynamodbav:"Created,omitempty"`
	FuncArn  string `dynamodbav:"FuncArn,omitempty"`
}

var ErrNotFound = errors.New("not found")

type DB interface {
	Get(ctx context.Context, clientId string) (*Record, error)
	Scan(ctx context.Context, f func(*Record)) error
	Put(ctx context.Context, record *Record) error
}
