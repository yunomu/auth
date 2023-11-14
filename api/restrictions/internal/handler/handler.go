package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/userlist"
	apipb "github.com/yunomu/auth/proto/api"
)

type Request events.APIGatewayV2HTTPRequest
type Response events.APIGatewayV2HTTPResponse

type Handler struct {
	userlistDB  userlist.DB
	marshaler   proto.MarshalOptions
	unmarshaler proto.UnmarshalOptions

	handlers map[string]func(context.Context, *Request) (proto.Message, error)

	logger Logger
}

func (h *Handler) list(ctx context.Context, req *Request) (proto.Message, error) {
	var users []*apipb.User
	if err := h.userlistDB.Scan(ctx, func(user *userlist.User, ts int64) {
		users = append(users, &apipb.User{
			Email:    user.Name,
			AppCodes: user.AppCodes,
			Version:  ts,
		})
	}); err != nil {
		h.logger.Error(err, "userlist.DB.Scan", req)
		return nil, err
	}

	return &apipb.RestrictionsResponse{
		Users: users,
	}, nil
}

func (h *Handler) post(ctx context.Context, req *Request) (proto.Message, error) {
	var user apipb.User
	if err := h.unmarshaler.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, err
	}

	version, err := h.userlistDB.Put(ctx, &userlist.User{
		Name:     user.Email,
		AppCodes: user.AppCodes,
	})
	if err != nil {
		return nil, err
	}
	user.Version = version

	return &apipb.RestrictionResponse{
		User: &user,
	}, nil
}

func (h *Handler) get(ctx context.Context, req *Request) (proto.Message, error) {
	return nil, errors.New("not implement")
}

func (h *Handler) put(ctx context.Context, req *Request) (proto.Message, error) {
	return nil, errors.New("not implement")
}

func (h *Handler) delete(ctx context.Context, req *Request) (proto.Message, error) {
	return nil, errors.New("not implement")
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
		marshaler:  proto.MarshalOptions{},
		unmarshaler: proto.UnmarshalOptions{
			DiscardUnknown: true,
		},

		logger: &defaultLogger{},
	}
	for _, f := range options {
		f(h)
	}

	h.handlers = map[string]func(context.Context, *Request) (proto.Message, error){
		"GET /v1/restrictions":            h.list,
		"POST /v1/restrictions":           h.post,
		"GET /v1/restrictions/{email}":    h.get,
		"PUT /v1/restrictions/{email}":    h.put,
		"DELETE /v1/restrictions/{email}": h.delete,
	}

	return h
}

func (h *Handler) Serve(ctx context.Context, req *Request) (*Response, error) {
	handler, ok := h.handlers[req.RouteKey]
	if !ok {
		h.logger.Info("not found", req)
		return &Response{
			StatusCode: 404,
			Body:       "NotFound",
		}, nil
	}

	res, err := handler(ctx, req)
	if err != nil {
		h.logger.Error(err, "handler error", req)
		return nil, err
	}

	var buf strings.Builder
	if res != nil {
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(res); err != nil {
			h.logger.Error(err, "json.Encoder.Encode", req, "response", res)
			return nil, err
		}
	}

	return &Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: buf.String(),
	}, nil
}
