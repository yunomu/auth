package userlist

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestGet(t *testing.T) {
	db := NewS3(&S3Mock{
		GetObjectFn: func(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			return &s3.GetObjectOutput{
				Body: io.NopCloser(strings.NewReader("abc@example.com\ntest@example.com")),
			}, nil
		},
	}, "bucket", "key")

	ctx := context.Background()
	exp := "test@example.com"
	res, err := db.Get(ctx, exp)
	if err != nil {
		t.Fatalf("unexpected error in db.Get: %v", err)
	}

	if res.Name != exp {
		t.Errorf("Name is mismatch: exp=`%v` act=`%v`", exp, res.Name)
	}
}

func TestGet_NoFile(t *testing.T) {
	db := NewS3(&S3Mock{
		GetObjectFn: func(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			return nil, &mockError{
				code: s3.ErrCodeNoSuchKey,
			}
		},
	}, "bucket", "key")

	ctx := context.Background()
	exp := "test@example.com"
	res, err := db.Get(ctx, exp)
	if err != ErrNoSuchUser {
		t.Logf("res=%v", res)
		t.Fatalf("unexpected error in db.Get: %v", err)
	}
}

func TestPut(t *testing.T) {
	db := NewS3(&S3Mock{
		GetObjectFn: func(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			return &s3.GetObjectOutput{
				Body: io.NopCloser(strings.NewReader("abc@example.com\ndef@example.com\n")),
			}, nil
		},
		PutObjectFn: func(in *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			if aws.StringValue(in.Bucket) != "bucket" {
				return nil, errors.New("bucket mismatch")
			}

			if aws.StringValue(in.Key) != "key" {
				return nil, errors.New("key mismatch")
			}

			r := csv.NewReader(in.Body)
			exp := map[string]bool{
				"abc@example.com":  false,
				"def@example.com":  false,
				"test@example.com": false,
			}
			for {
				rec, err := r.Read()
				if err == io.EOF {
					break
				} else if err != nil {
					return nil, err
				}

				if len(rec) != 1 {
					t.Errorf("body mismatch: exp=`%v` act=`%v`", exp, rec)
					return nil, errors.New("error")
				}

				if s, ok := exp[rec[0]]; !ok {
					t.Errorf("record is not found: rec=`%v`", s)
				}
				exp[rec[0]] = true
			}

			for k, v := range exp {
				if !v {
					t.Errorf("`%v` is not exist", k)
				}
			}

			return &s3.PutObjectOutput{}, nil
		},
	}, "bucket", "key")

	ctx := context.Background()
	if err := db.Put(ctx, &User{
		Name: "test@example.com",
	}); err != nil {
		t.Fatalf("PutError: %v", err)
	}
}

func TestPut_NoFile(t *testing.T) {
	exp := "test@example.com"
	db := NewS3(&S3Mock{
		GetObjectFn: func(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			return nil, &mockError{
				code: s3.ErrCodeNoSuchKey,
			}
		},
		PutObjectFn: func(in *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			if aws.StringValue(in.Bucket) != "bucket" {
				return nil, errors.New("bucket mismatch")
			}

			if aws.StringValue(in.Key) != "key" {
				return nil, errors.New("key mismatch")
			}

			r := csv.NewReader(in.Body)
			rec, err := r.Read()
			if err != nil {
				return nil, err
			}

			if len(rec) != 1 && rec[0] != exp {
				t.Errorf("body mismatch: exp=`%v` act=`%v`", exp, rec)
				return nil, errors.New("error")
			}

			rec2, err := r.Read()
			if err != io.EOF {
				t.Errorf("unexpected token: err=%v record=%v", err, rec2)
				return nil, errors.New("unknown field")
			}

			return &s3.PutObjectOutput{}, nil
		},
	}, "bucket", "key")

	ctx := context.Background()
	if err := db.Put(ctx, &User{
		Name: exp,
	}); err != nil {
		t.Fatalf("PutError: %v", err)
	}
}
