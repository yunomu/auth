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

func (h *handlerLogger) GetProductError(err error, clientId string, req *handler.Request) {
	h.logger.Error("GetProductError",
		"err", err,
		"clientId", clientId,
		"req", req,
	)
}

type productdbLogger struct {
	logger *slog.Logger
}

func (l *productdbLogger) GetError(err error, table string, key map[string]*dynamodb.AttributeValue) {
	l.logger.Debug("GetError",
		"err", err,
		"table", table,
		"key", key,
	)
}

func (l *productdbLogger) GetMarshalMapError(err error, clientId string) {
	l.logger.Debug("GetError in MarhalMap",
		"err", err,
		"clientId", clientId,
	)
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	productTable := os.Getenv("PRODUCT_TABLE")
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
			productdb.SetDynamoDBLogger(&productdbLogger{
				logger: logger.With("module", "productdb"),
			}),
		),
		handler.SetLogger(&handlerLogger{
			logger: logger.With("module", "handler"),
		}),
	)

	lambda.StartWithContext(ctx, h.Serve)
}
