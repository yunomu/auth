package preauthfunc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

type Lambda struct {
	client lambdaiface.LambdaAPI
}

func NewLambda(client lambdaiface.LambdaAPI) *Lambda {
	return &Lambda{
		client: client,
	}
}

func (l *Lambda) PreAuthentication(
	ctx context.Context,
	funcArn string,
	userName string,
) error {
	bs, err := json.Marshal(&Payload{
		UserName: userName,
	})
	if err != nil {
		return err
	}

	out, err := l.client.InvokeWithContext(ctx, &lambda.InvokeInput{
		FunctionName: aws.String(funcArn),
		Payload:      bs,
	})
	if err != nil {
		return err
	} else if out.FunctionError != nil {
		return errors.New(string(out.Payload))
	}

	return nil
}
