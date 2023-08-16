package main

import (
	"context"
	"os"

	"go.uber.org/zap"

	lambdahandler "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"

	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/preauthfunc"

	"github.com/yunomu/auth/lambda/preauth/internal/handler"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
}

type appLogger struct{}

func (a *appLogger) Error(err error, msg string) {
	logger.Error(msg, zap.Error(err))
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	productTable := os.Getenv("PRODUCT_TABLE")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Fatal("NewSession", zap.Error(err))
	}

	h := handler.NewHandler(
		productdb.NewDynamoDB(
			dynamodb.New(sess),
			productTable,
		),
		preauthfunc.NewLambda(
			lambda.New(sess),
		),
		handler.SetLogger(&appLogger{}),
	)

	lambdahandler.StartWithContext(ctx, h.Serve)
}
