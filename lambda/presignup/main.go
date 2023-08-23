package main

import (
	"context"
	"errors"
	"os"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/yunomu/auth/lambda/presignup/internal/handler"
	"github.com/yunomu/auth/lib/userlist"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
}

var errNoSuchKey = errors.New("no such key")

func loadWhiteList(ctx context.Context, client *s3.S3, bucket, whiteListKey string) ([]*userlist.User, error) {
	out, err := client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(whiteListKey),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeNoSuchKey {
				return nil, errNoSuchKey
			}
		}
		return nil, err
	}
	defer out.Body.Close()

	return userlist.Load(out.Body)
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	bucket := os.Getenv("BUCKET")
	whiteListKey := os.Getenv("WHITE_LIST_KEY")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Fatal("NewSession", zap.Error(err))
	}

	whiteList, err := loadWhiteList(ctx, s3.New(sess), bucket, whiteListKey)
	if err == errNoSuchKey {
		whiteList = []*userlist.User{}
	} else if err != nil {
		logger.Fatal("load white list",
			zap.Error(err),
			zap.String("BUCKET", bucket),
			zap.String("WHITE_LIST_KEY", whiteListKey),
		)
	}

	h := handler.NewHandler(
		whiteList,
	)

	lambda.StartWithContext(ctx, h.Serve)
}
