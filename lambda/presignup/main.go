package main

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

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

var errNoSuchKey = errors.New("no such key")

func loadWhiteList(ctx context.Context, client *s3.S3, whiteListObject string) ([]string, error) {
	ss := strings.SplitN(whiteListObject, "/", 2)
	if len(ss) != 2 {
		return nil, errors.New("Invalid Object URI")
	}
	bucket, key := ss[0], ss[1]

	out, err := client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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

	var ret []string
	r := csv.NewReader(out.Body)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(record) == 0 {
			continue
		}

		ret = append(ret, record[0])
	}

	return ret, nil
}

func main() {
	ctx := context.Background()

	region := os.Getenv("REGION")
	whiteListObject := os.Getenv("WHITE_LIST_OBJECT")

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		logger.Fatal("NewSession", zap.Error(err))
	}

	whiteList, err := loadWhiteList(ctx, s3.New(sess), whiteListObject)
	if err == errNoSuchKey {
		whiteList = []string{}
	} else if err != nil {
		logger.Fatal("load white list",
			zap.Error(err),
			zap.String("WHITE_LIST_OBJECT", whiteListObject),
		)
	}

	h := handler.NewHandler(
		whiteList,
	)

	lambda.StartWithContext(ctx, h.Serve)
}
