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

	"github.com/yunomu/auth/lambda/presignup/internal/handler"
	"github.com/yunomu/auth/lib/db/productdb"
	"github.com/yunomu/auth/lib/db/userlist"
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

func (h *handlerLogger) Signup(user *userlist.User) {
	h.logger.Info("SignUp",
		"user", user,
	)
}

func (h *handlerLogger) Info(msg string, email string, appCode string) {
	h.logger.Info(msg,
		"email", email,
		"appCode", appCode,
	)
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	productTable := os.Getenv("RESTRICTION_TABLE")
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
		productdb.NewDynamoDB(
			dynamodb.New(sess),
			productTable,
		),
		handler.SetLogger(&handlerLogger{
			logger: logger.With("module", "handler"),
		}),
	)

	lambda.StartWithContext(ctx, h.Serve)
}
