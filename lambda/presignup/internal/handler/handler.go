package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Handler struct {
	cognito    *cognitoidentityprovider.CognitoIdentityProvider
	userPoolId string
}

func NewHandler(
	cognito *cognitoidentityprovider.CognitoIdentityProvider,
	userPoolId string,
) *Handler {
	return &Handler{
		cognito:    cognito,
		userPoolId: userPoolId,
	}
}

type Request events.CognitoEventUserPoolsPreSignup

func (h *Handler) Serve(ctx context.Context, req *Request) (*Request, error) {
	req.Response.AutoConfirmUser = true // XXX: confirm all
	// TODO
	return req, nil
}
