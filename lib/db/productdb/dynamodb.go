package productdb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamoDB struct {
	client    dynamodbiface.DynamoDBAPI
	tableName string
}

var _ DB = (*DynamoDB)(nil)

func NewDynamoDB(
	client dynamodbiface.DynamoDBAPI,
	tableName string,
) *DynamoDB {
	return &DynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d *DynamoDB) Get(ctx context.Context, clientId string) (*Record, error) {
	key, err := dynamodbattribute.MarshalMap(&Record{
		ClientId: clientId,
	})
	if err != nil {
		return nil, err
	}

	out, err := d.client.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key:       key,
	})
	if err != nil {
		return nil, err
	} else if len(out.Item) == 0 {
		return nil, ErrNotFound
	}

	var ret Record
	if err := dynamodbattribute.UnmarshalMap(out.Item, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DynamoDB) Scan(ctx context.Context, f func(*Record)) error {
	var rerr error
	if err := d.client.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(d.tableName),
	}, func(out *dynamodb.ScanOutput, last bool) bool {
		select {
		case <-ctx.Done():
			rerr = ctx.Err()
			return false
		default:
		}

		var recs []Record
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			rerr = err
			return false
		}

		for _, rec := range recs {
			f(&rec)
		}

		return true
	}); rerr != nil {
		return rerr
	} else if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) Put(ctx context.Context, record *Record) error {
	av, err := dynamodbattribute.MarshalMap(record)
	if err != nil {
		return err
	}

	if _, err := d.client.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item:      av,
	}); err != nil {
		return err
	}

	return nil
}
