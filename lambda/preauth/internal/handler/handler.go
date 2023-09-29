package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"golang.org/x/sync/errgroup"

	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
	"github.com/yunomu/auth/lib/preauthfunc"
)

type Handler struct {
	productDB   productdb.DB
	userlistDB  userlist.DB
	preAuthFunc preauthfunc.Func

	logger Logger
}

type Option func(*Handler)

func SetLogger(l Logger) Option {
	return func(h *Handler) {
		if l != nil {
			h.logger = &defaultLogger{}
		} else {
			h.logger = l
		}
	}
}

func NewHandler(
	productDB productdb.DB,
	userlistDB userlist.DB,
	preAuthFunc preauthfunc.Func,
	opts ...Option,
) *Handler {
	h := &Handler{
		productDB:   productDB,
		userlistDB:  userlistDB,
		preAuthFunc: preAuthFunc,

		logger: &defaultLogger{},
	}
	for _, f := range opts {
		f(h)
	}

	return h
}

func containsAppCode(appCode string, appCodes []string) bool {
	for _, c := range appCodes {
		if c == appCode {
			return true
		}
	}
	return false
}

type Request events.CognitoEventUserPoolsPreAuthentication

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	g, ctx := errgroup.WithContext(ctx)

	var product *productdb.Product
	g.Go(func() error {
		rec, _, err := h.productDB.Get(ctx, req.CallerContext.ClientID)
		if err != nil {
			h.logger.Error(err, "productDB.Get")
			return err
		}

		product = rec

		return nil
	})

	var user *userlist.User
	g.Go(func() error {
		u, _, err := h.userlistDB.Get(ctx, req.UserName)
		if err != nil {
			return err
		}

		user = u

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if !containsAppCode(product.AppCode, user.AppCodes) {
		h.logger.Info("the user does not have permissions for this app", req, product, user)
		return req, nil
	}

	if product.FuncArn == "" {
		return req, nil
	}

	if err := h.preAuthFunc.PreAuthentication(ctx, product.FuncArn, req.UserName); err != nil {
		h.logger.Error(err, "preAuthFunc.PreAuthentication")
		return nil, err
	}

	return req, nil
}
