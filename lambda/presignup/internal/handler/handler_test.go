package handler

import (
	"testing"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
)

type testUserlistDB struct {
	GetFn func(context.Context, string) (*userlist.User, int64, error)
}

func (t *testUserlistDB) Get(ctx context.Context, name string) (*userlist.User, int64, error) {
	if f := t.GetFn; f != nil {
		return f(ctx, name)
	}
	panic("not assigned")
}

func (t *testUserlistDB) Scan(ctx context.Context, f func(*userlist.User, int64)) error {
	panic("not implemented") // TODO: Implement
}

func (t *testUserlistDB) Put(ctx context.Context, user *userlist.User) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testUserlistDB) Update(ctx context.Context, user *userlist.User, timestamp int64) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testUserlistDB) Delete(ctx context.Context, name string, version int64) error {
	panic("not implemented") // TODO: Implement
}

type testProductDB struct {
	GetFn func(context.Context, string) (*productdb.Product, int64, error)
}

func (t *testProductDB) Get(ctx context.Context, clientId string) (*productdb.Product, int64, error) {
	if f := t.GetFn; f != nil {
		return f(ctx, clientId)
	}
	panic("not assigned")
}

func (t *testProductDB) Scan(ctx context.Context, f func(*productdb.Product, int64)) error {
	panic("not implemented") // TODO: Implement
}

func (t *testProductDB) Put(ctx context.Context, record *productdb.Product) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testProductDB) Update(ctx context.Context, record *productdb.Product, version int64) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testProductDB) Delete(ctx context.Context, clientId string, version int64) error {
	panic("not implemented") // TODO: Implement
}

func TestServe_HasNotPermission(t *testing.T) {
	h := NewHandler(
		&testUserlistDB{
			GetFn: func(_ context.Context, name string) (*userlist.User, int64, error) {
				return &userlist.User{
					Name:     "test@example.com",
					AppCodes: []string{"aaa", "bbb"},
				}, 0, nil
			},
		},
		&testProductDB{
			GetFn: func(_ context.Context, clientId string) (*productdb.Product, int64, error) {
				return &productdb.Product{
					AppCode: "testapp",
				}, 0, nil
			},
		},
	)

	ctx := context.Background()
	res, err := h.Serve(ctx, &Request{
		CognitoEventUserPoolsHeader: events.CognitoEventUserPoolsHeader{
			CallerContext: events.CognitoEventUserPoolsCallerContext{},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Response.AutoConfirmUser {
		t.Errorf("autoConfirmUser is true")
	}
}
