package userlist

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamoDB struct {
	api   dynamodbiface.DynamoDBAPI
	table string
}

func NewDynamoDB(
	api dynamodbiface.DynamoDBAPI,
	table string,
) *DynamoDB {
	return &DynamoDB{
		api:   api,
		table: table,
	}
}

type record struct {
	Email    string   `dynamodbav:"Email"`
	AppCodes []string `dynamodbav:"AppCodes"`
}

func (db *DynamoDB) Get(ctx context.Context, email string) (*User, error) {
	key, err := dynamodbattribute.MarshalMap(&record{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	out, err := db.api.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.table),
		Key:       key,
	})
	if err != nil {
		return nil, err
	} else if len(out.Item) == 0 {
		return nil, ErrNoSuchUser
	}

	var rec record
	if err := dynamodbattribute.UnmarshalMap(out.Item, &rec); err != nil {
		return nil, err
	}

	return &User{
		Name: rec.Email,
	}, nil
}

func (db *DynamoDB) Scan(ctx context.Context, f func(*User)) error {
	var rerr error
	if err := db.api.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(db.table),
	}, func(out *dynamodb.ScanOutput, last bool) bool {
		select {
		case <-ctx.Done():
			rerr = ctx.Err()
			return false
		default:
		}

		var recs []record
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			rerr = err
			return false
		}

		for _, rec := range recs {
			f(&User{
				Name: rec.Email,
			})
		}

		return true
	}); rerr != nil {
		return rerr
	} else if err != nil {
		return err
	}

	return nil
}

func (db *DynamoDB) Put(ctx context.Context, user *User) error {
	av, err := dynamodbattribute.MarshalMap(&record{
		Email: user.Name,
	})
	if err != nil {
		return err
	}

	if _, err := db.api.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      av,
	}); err != nil {
		return err
	}

	return nil
}
