package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/productdb"
)

type Handler struct {
	db productdb.DB
}

func NewHandler(
	db productdb.DB,
) *Handler {
	return &Handler{
		db: db,
	}
}

type Request events.CognitoEventUserPoolsPreAuthentication

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	// TODO
	return req, nil
}
