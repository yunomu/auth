package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	lambdahandler "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"

	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
	"github.com/yunomu/auth/lib/preauthfunc"

	"github.com/yunomu/auth/lambda/preauth/internal/handler"
)

var (
	logger *slog.Logger

	debug = flag.Bool("g", false, "Debug")
)

func init() {
	levelVar := new(slog.LevelVar)
	if *debug {
		levelVar.Set(slog.LevelDebug)
	}

	logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: levelVar,
	}))
}

type appLogger struct {
	logger *slog.Logger
}

func (a *appLogger) Error(err error, msg string) {
	a.logger.Error(msg, "err", err)
}

func (a *appLogger) Info(msg string, req *handler.Request, product *productdb.Record, user *userlist.User) {
	a.logger.Info(msg,
		"request", req,
		"product", product,
		"user", user,
	)
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	productTable := os.Getenv("PRODUCT_TABLE")
	restrictionTable := os.Getenv("RESTRICTION_TABLE")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Error("NewSession", "err", err)
		return
	}

	h := handler.NewHandler(
		productdb.NewDynamoDB(
			dynamodb.New(sess),
			productTable,
		),
		userlist.NewDynamoDB(
			dynamodb.New(sess),
			restrictionTable,
		),
		preauthfunc.NewLambda(
			lambda.New(sess),
		),
		handler.SetLogger(&appLogger{
			logger: logger.With("module", "handler"),
		}),
	)

	lambdahandler.StartWithContext(ctx, h.Serve)
}
