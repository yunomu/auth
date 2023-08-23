package handler

import (
	"context"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yunomu/auth/lib/userlist"
)

type Handler struct {
	whiteList []*userlist.User
}

type userSlice []*userlist.User

func (u userSlice) Len() int {
	return len(u)
}

func (u userSlice) Less(i int, j int) bool {
	return strings.Compare(u[i].Name, u[j].Name) < 0
}

func (u userSlice) Swap(i int, j int) {
	u[i], u[j] = u[j], u[i]
}

func NewHandler(
	whiteList []*userlist.User,
) *Handler {
	sort.Sort(userSlice(whiteList))
	return &Handler{
		whiteList: whiteList,
	}
}

type Request events.CognitoEventUserPoolsPreSignup

func (h *Handler) containsEmail(email string) bool {
	return sort.Search(len(h.whiteList), func(i int) bool {
		return h.whiteList[i].Name == email
	}) != -1
}

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	req.Response.AutoConfirmUser = h.containsEmail(req.Request.UserAttributes["email"])

	return req, nil
}
