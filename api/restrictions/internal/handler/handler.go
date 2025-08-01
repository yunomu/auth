package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/userlist"
	apipb "github.com/yunomu/auth/proto/api"
)

type Request events.APIGatewayV2HTTPRequest
type Response events.APIGatewayV2HTTPResponse

type BadRequest struct {
	Message string `json:"message"`
}

func (e *BadRequest) Code() int {
	return 400
}

func (e *BadRequest) Error() string {
	var buf strings.Builder
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(e); err != nil {
		return e.Message
	}

	return buf.String()
}

type Handler struct {
	userlistDB  userlist.DB
	marshaler   protojson.MarshalOptions
	unmarshaler protojson.UnmarshalOptions

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

	version, err := h.userlistDB.Update(ctx, &userlist.User{
		Name:     user.Email,
		AppCodes: user.AppCodes,
	}, 0)
	if err != nil {
		if err == userlist.ErrOptimisticLock {
			return nil, &BadRequest{
				Message: "user already exists",
			}
		}
		return nil, err
	}
	user.Version = version

	return &apipb.RestrictionResponse{
		User: &user,
	}, nil
}

func (h *Handler) get(ctx context.Context, req *Request) (proto.Message, error) {
	email, ok := req.PathParameters["email"]
	if !ok {
		return nil, errors.New("PathParameters[`email`] not found")
	}

	user, version, err := h.userlistDB.Get(ctx, email)
	if err != nil {
		if err == userlist.ErrNoSuchUser {
			return nil, &BadRequest{
				Message: "user not found",
			}
		}
		return nil, err
	}

	return &apipb.RestrictionResponse{
		User: &apipb.User{
			Email:    user.Name,
			AppCodes: user.AppCodes,
			Version:  version,
		},
	}, nil
}

func (h *Handler) put(ctx context.Context, req *Request) (proto.Message, error) {
	email, ok := req.PathParameters["email"]
	if !ok {
		return nil, errors.New("PathParameters[`email`] not found")
	}

	var user apipb.User
	if err := h.unmarshaler.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, err
	}

	if email != user.Email {
		return nil, &BadRequest{
			Message: "object and path are mismatch",
		}
	}

	if _, err := h.userlistDB.Update(ctx, &userlist.User{
		Name:     user.Email,
		AppCodes: user.AppCodes,
	}, user.Version); err != nil {
		if err == userlist.ErrOptimisticLock {
			return nil, &BadRequest{
				Message: "lock error",
			}
		}
		return nil, err
	}

	return nil, nil
}

func (h *Handler) delete(ctx context.Context, req *Request) (proto.Message, error) {
	email, ok := req.PathParameters["email"]
	if !ok {
		return nil, errors.New("PathParameters[`email`] not found")
	}

	versionStr, ok := req.QueryStringParameters["version"]
	if !ok {
		return nil, &BadRequest{
			Message: "version parameter is not found",
		}
	}

	version, err := strconv.ParseInt(versionStr, 10, 64)
	if err != nil {
		return nil, &BadRequest{
			Message: "version is not a number",
		}
	}

	if err := h.userlistDB.Delete(ctx, email, version); err != nil {
		if err == userlist.ErrOptimisticLock {
			return nil, &BadRequest{
				Message: "lock error",
			}
		}
		return nil, err
	}

	return nil, nil
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
		marshaler: protojson.MarshalOptions{
			EmitUnpopulated: false,
		},
		unmarshaler: protojson.UnmarshalOptions{
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
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: "NotFound",
		}, nil
	}

	res, err := handler(ctx, req)
	if err != nil {
		if badReq, ok := err.(*BadRequest); ok {
			return &Response{
				StatusCode: badReq.Code(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: badReq.Error(),
			}, nil
		}
		h.logger.Error(err, "handler error", req)
		return nil, err
	}

	var buf strings.Builder
	if res != nil {
		bs, err := h.marshaler.Marshal(res)
		if err != nil {
			h.logger.Error(err, "proto.Marshal", req, "response", res)
			return nil, err
		}
		if _, err := buf.Write(bs); err != nil {
			h.logger.Error(err, "strings.Builder.Write", req, "response", res, "bytes", bs)
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
