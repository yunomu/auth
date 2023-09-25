package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
)

type Handler struct {
	userlistDB userlist.DB
	productDB  productdb.DB

	logger Logger
}

type Option func(*Handler)

func SetLogger(l Logger) Option {
	return func(h *Handler) {
		if l == nil {
			h.logger = &defaultLogger{}
		} else {
			h.logger = l
		}
	}
}

func NewHandler(
	userlistDB userlist.DB,
	productDB productdb.DB,
	opts ...Option,
) *Handler {
	h := &Handler{
		userlistDB: userlistDB,
		productDB:  productDB,

		logger: &defaultLogger{},
	}
	for _, f := range opts {
		f(h)
	}

	return h
}

type Request events.CognitoEventUserPoolsPreSignup

func containsAppCode(code string, codes []string) bool {
	for _, c := range codes {
		if c == code {
			return true
		}
	}
	return false
}

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	user, err := h.userlistDB.Get(ctx, req.UserName)
	if err == userlist.ErrNoSuchUser {
		h.logger.Info("no such user", req.UserName, "(null)")
		return req, nil
	} else if err != nil {
		return nil, err
	}

	clientID := req.CallerContext.ClientID
	product, err := h.productDB.Get(ctx, clientID)
	if err != nil {
		return nil, err
	}

	if !containsAppCode(product.AppCode, user.AppCodes) {
		h.logger.Info("the user does not have permissions for this app", req.UserName, product.AppCode)
		return req, nil
	}

	h.logger.Signup(user)
	req.Response.AutoConfirmUser = true

	return req, nil
}
