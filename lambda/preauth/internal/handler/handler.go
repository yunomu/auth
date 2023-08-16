package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/preauthfunc"
)

type Handler struct {
	db          productdb.DB
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
	db productdb.DB,
	preAuthFunc preauthfunc.Func,
	opts ...Option,
) *Handler {
	h := &Handler{
		db:          db,
		preAuthFunc: preAuthFunc,

		logger: &defaultLogger{},
	}
	for _, f := range opts {
		f(h)
	}

	return h
}

type Request events.CognitoEventUserPoolsPreAuthentication

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	rec, err := h.db.Get(ctx, req.CallerContext.ClientID)
	if err != nil {
		h.logger.Error(err, "db.Get")
		return nil, err
	}

	if rec.FuncArn == "" {
		return req, nil
	}

	if err := h.preAuthFunc.PreAuthentication(ctx, rec.FuncArn, req.UserName); err != nil {
		h.logger.Error(err, "preAuthFunc.PreAuthentication")
		return nil, err
	}

	return req, nil
}
