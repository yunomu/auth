package userlist

import (
	"bytes"
	"context"
	"encoding/csv"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type S3 struct {
	api         s3iface.S3API
	bucket, key string
}

func NewS3(api s3iface.S3API, bucket, key string) *S3 {
	return &S3{
		api:    api,
		bucket: bucket,
		key:    key,
	}
}

func load(ctx context.Context, in io.Reader, f func(*User)) error {
	r := csv.NewReader(in)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(record) == 0 {
			continue
		}

		f(&User{
			Name: record[0],
		})
	}

	return nil
}

func (db *S3) Scan(ctx context.Context, f func(*User)) error {
	out, err := db.api.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(db.bucket),
		Key:    aws.String(db.key),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case s3.ErrCodeNoSuchKey:
				return nil
			default:
			}
		}
		return err
	}
	defer out.Body.Close()

	return load(ctx, out.Body, func(u *User) {
		f(u)
	})
}

func (db *S3) Get(ctx context.Context, name string) (*User, error) {
	ctx, cancel := context.WithCancel(ctx)

	var ret *User
	if err := db.Scan(ctx, func(u *User) {
		if u.Name == name {
			ret = u
			cancel()
		}
	}); err == context.Canceled {
		// do nothing
	} else if err != nil {
		return nil, err
	}

	if ret == nil {
		return nil, ErrNoSuchUser
	}

	return ret, nil
}

func (db *S3) getAll(ctx context.Context) ([]*User, error) {
	var ret []*User
	if err := db.Scan(ctx, func(u *User) {
		ret = append(ret, u)
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (db *S3) Put(ctx context.Context, user *User) error {
	users, err := db.getAll(ctx)
	if err != nil {
		return err
	}

	var idx = -1
	for i, u := range users {
		if u.Name == user.Name {
			idx = i
			break
		}
	}
	if idx == -1 {
		users = append(users, user)
	} else {
		users[idx] = user
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	for _, line := range users {
		if err := w.Write([]string{
			line.Name,
		}); err != nil {
			return err
		}
	}
	w.Flush()

	if _, err := db.api.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(db.bucket),
		Key:    aws.String(db.key),
		Body:   bytes.NewReader(buf.Bytes()),
	}); err != nil {
		return err
	}

	return nil
}
