package userlist

import (
	"bytes"
	"context"
	"encoding/csv"
	"io"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func Load(in io.Reader) ([]*User, error) {
	var ret []*User
	r := csv.NewReader(in)
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

		ret = append(ret, &User{
			Name: record[0],
		})
	}

	return ret, nil
}

type S3 struct {
	api         s3iface.S3API
	bucket, key string

	cache []*User
	dirty bool
}

func NewS3(api s3iface.S3API, bucket, key string) *S3 {
	return &S3{
		api:    api,
		bucket: bucket,
		key:    key,
		dirty:  true,
	}
}

type userSlice []*User

func (u userSlice) Len() int {
	return len(u)
}

func (u userSlice) Less(i int, j int) bool {
	return strings.Compare(u[i].Name, u[j].Name) < 0
}

func (u userSlice) Swap(i int, j int) {
	u[i], u[j] = u[j], u[i]
}

func (db *S3) reload(ctx context.Context) error {
	if !db.dirty {
		return nil
	}

	out, err := db.api.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(db.bucket),
		Key:    aws.String(db.key),
	})
	if err != nil {
		return nil
	}
	defer out.Body.Close()

	users, err := Load(out.Body)
	if err != nil {
		return nil
	}

	sort.Sort(userSlice(users))
	db.cache = users
	db.dirty = false

	return nil
}

func (db *S3) PreLoad(ctx context.Context) error {
	db.dirty = true
	return db.reload(ctx)
}

func (db *S3) get(ctx context.Context, name string) (int, error) {
	if err := db.reload(ctx); err != nil {
		return -1, err
	}

	idx := sort.Search(len(db.cache), func(i int) bool {
		return db.cache[i].Name == name
	})

	if idx < len(db.cache) && db.cache[idx].Name == name {
		return idx, nil
	} else {
		return -1, ErrNoSuchUser
	}
}

func (db *S3) Get(ctx context.Context, name string) (*User, error) {
	idx, err := db.get(ctx, name)
	if err != nil {
		return nil, err
	}

	return db.cache[idx], nil
}

func (db *S3) Scan(ctx context.Context, f func(*User)) error {
	if err := db.reload(ctx); err != nil {
		return err
	}

	for _, u := range db.cache {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			f(u)
		}
	}

	return nil
}

func (db *S3) Put(ctx context.Context, user *User) error {
	if err := db.PreLoad(ctx); err != nil {
		return err
	}

	if idx, err := db.get(ctx, user.Name); err == ErrNoSuchUser {
		db.cache = append(db.cache, user)
	} else if err != nil {
		return err
	} else {
		db.cache[idx] = user
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	for _, line := range db.cache {
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
