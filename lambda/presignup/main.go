package main

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/yunomu/auth/lambda/presignup/internal/handler"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	userPoolId := os.Getenv("USER_POOL_ID")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Fatal("NewSession", zap.Error(err))
	}

	h := handler.NewHandler(
		cognitoidentityprovider.New(
			sess,
		),
		userPoolId,
	)

	lambda.StartWithContext(ctx, h.Serve)
}
