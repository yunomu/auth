package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/userlist"
)

type Handler struct {
	userlistDB userlist.DB

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
	options ...Option,
) *Handler {
	h := &Handler{
		userlistDB: userlistDB,

		logger: &defaultLogger{},
	}
	for _, f := range options {
		f(h)
	}

	return h
}

type Request events.APIGatewayV2HTTPRequest
type Response events.APIGatewayV2HTTPResponse

type User struct {
	Email    string   `json:"email"`
	AppCodes []string `json:"appcodes"`
	Version  int64    `json:"version"`
}

func (h *Handler) list(ctx context.Context, req *Request) (*Response, error) {
	var users []*User
	if err := h.userlistDB.Scan(ctx, func(user *userlist.User, ts int64) {
		users = append(users, &User{
			Email:    user.Name,
			AppCodes: user.AppCodes,
			Version:  ts,
		})
	}); err != nil {
		h.logger.Error(err, "userlist.DB.Scan", req)
		return nil, err
	}

	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {
		h.logger.Error(err, "json.Encoder.Encode", req, "users", users)
		return nil, err
	}

	return &Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: buf.String(),
	}, nil
}

func (h *Handler) Serve(ctx context.Context, req *Request) (*Response, error) {
	switch req.RouteKey {
	case "GET /v1/restrictions":
		return h.list(ctx, req)
	default:
		h.logger.Info("not found", req)
		return &Response{
			StatusCode: 404,
			Body:       "NotFound",
		}, nil
	}

	// not reached
	h.logger.Error(nil, "not reached", req)
	return &Response{
		StatusCode: 503,
		Body:       "InternalError",
	}, nil
}
