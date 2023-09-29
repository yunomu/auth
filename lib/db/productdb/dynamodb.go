package productdb

import (
	"context"
	"time"

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

type dynamodbRecord struct {
	ClientId  string `dynamodbav:"ClientId"`
	AppCode   string `dynamodbav:"AppCode,omitempty"`
	Timestamp int64  `dynamodbav:"Timestamp,omitempty"`
	FuncArn   string `dynamodbav:"FuncArn,omitempty"`
}

func toDynamoDB(p *Product, ts int64) *dynamodbRecord {
	return &dynamodbRecord{
		ClientId:  p.ClientId,
		AppCode:   p.AppCode,
		FuncArn:   p.FuncArn,
		Timestamp: ts,
	}
}

func fromDynamoDB(r *dynamodbRecord) *Product {
	return &Product{
		ClientId: r.ClientId,
		AppCode:  r.AppCode,
		FuncArn:  r.FuncArn,
	}
}

func NewDynamoDB(
	client dynamodbiface.DynamoDBAPI,
	tableName string,
) *DynamoDB {
	return &DynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d *DynamoDB) Get(ctx context.Context, clientId string) (*Product, int64, error) {
	key, err := dynamodbattribute.MarshalMap(&dynamodbRecord{
		ClientId: clientId,
	})
	if err != nil {
		return nil, -1, err
	}

	out, err := d.client.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key:       key,
	})
	if err != nil {
		return nil, -1, err
	} else if len(out.Item) == 0 {
		return nil, -1, ErrNotFound
	}

	var rec dynamodbRecord
	if err := dynamodbattribute.UnmarshalMap(out.Item, &rec); err != nil {
		return nil, -1, err
	}

	return fromDynamoDB(&rec), rec.Timestamp, nil
}

func (d *DynamoDB) Scan(ctx context.Context, f func(*Product, int64)) error {
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

		var recs []dynamodbRecord
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			rerr = err
			return false
		}

		for _, rec := range recs {
			f(fromDynamoDB(&rec), rec.Timestamp)
		}

		return true
	}); rerr != nil {
		return rerr
	} else if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) Put(ctx context.Context, product *Product) error {
	ts := time.Now().UnixMicro()

	av, err := dynamodbattribute.MarshalMap(toDynamoDB(product, ts))
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
