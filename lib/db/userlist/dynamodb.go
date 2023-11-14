package userlist

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
	Email     string   `dynamodbav:"Email"`
	AppCodes  []string `dynamodbav:"AppCodes,omitempty"`
	Timestamp int64    `dynamodbav:"Timestamp,omitempty"`
}

func (db *DynamoDB) Get(ctx context.Context, email string) (*User, int64, error) {
	key, err := dynamodbattribute.MarshalMap(&record{
		Email: email,
	})
	if err != nil {
		return nil, -1, err
	}

	out, err := db.api.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.table),
		Key:       key,
	})
	if err != nil {
		return nil, -1, err
	} else if len(out.Item) == 0 {
		return nil, -1, ErrNoSuchUser
	}

	var rec record
	if err := dynamodbattribute.UnmarshalMap(out.Item, &rec); err != nil {
		return nil, -1, err
	}

	return &User{
		Name:     rec.Email,
		AppCodes: rec.AppCodes,
	}, rec.Timestamp, nil
}

func (db *DynamoDB) Scan(ctx context.Context, f func(*User, int64)) error {
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
				Name:     rec.Email,
				AppCodes: rec.AppCodes,
			}, rec.Timestamp)
		}

		return true
	}); rerr != nil {
		return rerr
	} else if err != nil {
		return err
	}

	return nil
}

func (db *DynamoDB) Put(ctx context.Context, user *User) (int64, error) {
	timestamp := time.Now().UnixMicro()
	av, err := dynamodbattribute.MarshalMap(&record{
		Email:     user.Name,
		AppCodes:  user.AppCodes,
		Timestamp: timestamp,
	})
	if err != nil {
		return 0, err
	}

	if _, err := db.api.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      av,
	}); err != nil {
		return 0, err
	}

	return timestamp, nil
}
func (db *DynamoDB) Update(ctx context.Context, user *User, timestamp int64) error {
	expr, err := expression.NewBuilder().
		WithCondition(expression.Name("Timestamp").Equal(expression.Value(timestamp))).
		Build()
	if err != nil {
		return err
	}

	av, err := dynamodbattribute.MarshalMap(&record{
		Email:     user.Name,
		AppCodes:  user.AppCodes,
		Timestamp: time.Now().UnixMicro(),
	})
	if err != nil {
		return err
	}

	if _, err := db.api.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      av,

		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}); err != nil {
		return err
	}

	return nil
}
