package handler

import (
	"context"
	"sort"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	whiteList    []string
	whiteListLen int
}

func NewHandler(
	whiteList []string,
) *Handler {
	sort.StringSlice(whiteList).Sort()
	return &Handler{
		whiteList: whiteList,
	}
}

type Request events.CognitoEventUserPoolsPreSignup

func (h *Handler) containsEmail(email string) bool {
	return sort.SearchStrings(h.whiteList, email) != -1
}

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	req.Response.AutoConfirmUser = h.containsEmail(req.Request.UserAttributes["email"])

	return req, nil
}
