package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/productdb"
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
	productDB   productdb.DB
	marshaler   proto.MarshalOptions
	unmarshaler proto.UnmarshalOptions

	handlers map[string]func(context.Context, *Request) (proto.Message, error)

	logger Logger
}

func (h *Handler) list(ctx context.Context, req *Request) (proto.Message, error) {
	var products []*apipb.Product
	if err := h.productDB.Scan(ctx, func(p *productdb.Product, ts int64) {
		products = append(products, &apipb.Product{
			ClientId: p.ClientId,
			AppCode:  p.AppCode,
			FuncArn:  p.FuncArn,
			Version:  ts,
		})
	}); err != nil {
		h.logger.Error(err, "productdb.DB.Scan", req)
		return nil, err
	}

	return &apipb.ProductsResponse{
		Products: products,
	}, nil
}

func (h *Handler) post(ctx context.Context, req *Request) (proto.Message, error) {
	var p apipb.Product
	if err := h.unmarshaler.Unmarshal([]byte(req.Body), &p); err != nil {
		h.logger.Error(err, "Unmarshal", req)
		return nil, &BadRequest{
			Message: "unrecognized format",
		}
	}

	product := &productdb.Product{
		ClientId: p.ClientId,
		AppCode:  p.AppCode,
		FuncArn:  p.FuncArn,
	}
	version, err := h.productDB.Update(ctx, product, 0)
	if err != nil {
		if err == productdb.ErrOptimisticLock {
			return nil, &BadRequest{
				Message: "lock error",
			}
		}
		h.logger.Error(err, "productdb.DB.Update", req,
			"product", product,
		)
		return nil, err
	}

	p.Version = version
	return &apipb.ProductResponse{
		Product: &p,
	}, nil
}

func (h *Handler) get(ctx context.Context, req *Request) (proto.Message, error) {
	clientId, ok := req.PathParameters["clientId"]
	if !ok {
		return nil, errors.New("PathParameters[`clientId`] is not found")
	}

	product, version, err := h.productDB.Get(ctx, clientId)
	if err != nil {
		if err == productdb.ErrNotFound {
			return nil, &BadRequest{
				Message: "product not found",
			}
		}
		return nil, err
	}

	return &apipb.ProductResponse{
		Product: &apipb.Product{
			ClientId: product.ClientId,
			AppCode:  product.AppCode,
			FuncArn:  product.FuncArn,
			Version:  version,
		},
	}, nil
}

func (h *Handler) put(ctx context.Context, req *Request) (proto.Message, error) {
	clientId, ok := req.PathParameters["clintId"]
	if !ok {
		return nil, errors.New("PathParameters[`clientId`] is not found")
	}

	var p apipb.Product
	if err := h.unmarshaler.Unmarshal([]byte(req.Body), &p); err != nil {
		h.logger.Error(err, "Unmarshal", req)
		return nil, &BadRequest{
			Message: "unrecognized format",
		}
	}

	if clientId != p.ClientId {
		return nil, &BadRequest{
			Message: "invalid data",
		}
	}

	product := &productdb.Product{
		ClientId: p.ClientId,
		AppCode:  p.AppCode,
		FuncArn:  p.FuncArn,
	}
	if _, err := h.productDB.Update(ctx, product, p.Version); err != nil {
		if err == productdb.ErrOptimisticLock {
			return nil, &BadRequest{
				Message: "lock error",
			}
		}
		h.logger.Error(err, "productdb.DB.Update", req,
			"product", product,
		)
		return nil, err
	}

	return nil, nil
}

func (h *Handler) delete(ctx context.Context, req *Request) (proto.Message, error) {
	clientId, ok := req.PathParameters["clientId"]
	if !ok {
		return nil, errors.New("PathParameters[`clientId`] is not found")
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
			Message: "version is invalid format",
		}
	}

	if err := h.productDB.Delete(ctx, clientId, version); err != nil {
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
	productDB productdb.DB,
	options ...Option,
) *Handler {
	h := &Handler{
		productDB: productDB,
		marshaler: proto.MarshalOptions{},
		unmarshaler: proto.UnmarshalOptions{
			DiscardUnknown: true,
		},

		logger: &defaultLogger{},
	}
	for _, f := range options {
		f(h)
	}

	h.handlers = map[string]func(context.Context, *Request) (proto.Message, error){
		"GET /v1/products":               h.list,
		"POST /v1/products":              h.post,
		"GET /v1/products/{clientId}":    h.get,
		"PUT /v1/products/{clientId}":    h.put,
		"DELETE /v1/products/{clientId}": h.delete,
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
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(res); err != nil {
		h.logger.Error(err, "json.Encoder.Encode", req, "response", res)
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
