package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/userlist"
)

type Handler struct {
	userlistDB userlist.DB

	logger Logger
}

func NewHandler(
	userlistDB userlist.DB,
) *Handler {
	return &Handler{
		userlistDB: userlistDB,

		logger: &defaultLogger{},
	}
}

type Request events.CognitoEventUserPoolsPreSignup

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	var autoConfirm bool
	user, err := h.userlistDB.Get(ctx, req.Request.UserAttributes["email"])
	if err == userlist.ErrNoSuchUser {
		// do nothing
	} else if err != nil {
		return nil, err
	} else {
		autoConfirm = true
		h.logger.Signup(user)
	}

	req.Response.AutoConfirmUser = autoConfirm

	return req, nil
}
