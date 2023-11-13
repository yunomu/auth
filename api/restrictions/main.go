package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/yunomu/auth/lib/db/userlist"

	"github.com/yunomu/auth/api/restrictions/internal/handler"
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

type handlerLogger struct {
	logger *slog.Logger
}

func (h *handlerLogger) Error(err error, msg string, req *handler.Request, args ...interface{}) {
	h.logger.Error(msg, append([]interface{}{
		"error", err,
		"request", req,
	}, args...)...)
}

func (h *handlerLogger) Info(msg string, req *handler.Request, args ...interface{}) {
	h.logger.Info(msg, append([]interface{}{
		"request", req,
	}, args...)...)
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	restrictionTable := os.Getenv("RESTRICTION_TABLE")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Error("NewSession",
			"err", err,
			"region", region,
		)
		return
	}

	h := handler.NewHandler(
		userlist.NewDynamoDB(
			dynamodb.New(sess),
			restrictionTable,
		),
		handler.SetLogger(&handlerLogger{
			logger: logger.With("module", "handler"),
		}),
	)

	lambda.StartWithContext(ctx, h.Serve)
}
