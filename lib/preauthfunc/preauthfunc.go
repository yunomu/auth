package preauthfunc

import "context"

type Payload struct {
	UserName string `json:"username"`
}

type Func interface {
	PreAuthentication(_ context.Context, funcArn, userName string) error
}
