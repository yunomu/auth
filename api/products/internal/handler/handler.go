package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yunomu/auth/lib/db/productdb"
	apipb "github.com/yunomu/auth/proto/api"
)

type Request events.APIGatewayV2HTTPRequest
type Response events.APIGatewayV2HTTPResponse

type Handler struct {
	productDB productdb.DB

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
	return nil, errors.New("not implement")
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
	productDB productdb.DB,
	options ...Option,
) *Handler {
	h := &Handler{
		productDB: productDB,

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
