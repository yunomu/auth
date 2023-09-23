package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/yunomu/auth/lambda/presignup/internal/handler"
	"github.com/yunomu/auth/lib/userlist"
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

var errNoSuchKey = errors.New("no such key")

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	bucket := os.Getenv("BUCKET")
	whiteListKey := os.Getenv("WHITE_LIST_KEY")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Error("NewSession",
			"err", err,
			"region", region,
		)
		return
	}

	h := handler.NewHandler(
		userlist.NewS3(
			s3.New(sess),
			bucket,
			whiteListKey,
		),
	)

	lambda.StartWithContext(ctx, h.Serve)
}
