package handler

import (
	"context"

	"golang.org/x/sync/errgroup"

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
	g, ctx := errgroup.WithContext(ctx)

	var user *userlist.User
	g.Go(func() error {
		u, _, err := h.userlistDB.Get(ctx, req.Request.UserAttributes["email"])
		if err == userlist.ErrNoSuchUser {
			h.logger.Info("no such user", req.UserName, "(null)")
			return nil
		} else if err != nil {
			return err
		}

		user = u

		return nil
	})

	var product *productdb.Product
	g.Go(func() error {
		clientID := req.CallerContext.ClientID
		p, _, err := h.productDB.Get(ctx, clientID)
		if err != nil {
			h.logger.GetProductError(err, clientID, req)
			return err
		}

		product = p

		return nil
	})

	if err := g.Wait(); err != nil {
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
