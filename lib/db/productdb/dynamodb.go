package productdb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamoDBLogger interface {
	GetError(err error, table string, key map[string]*dynamodb.AttributeValue)
	GetMarshalMapError(err error, clientId string)
}

type DynamoDB struct {
	client    dynamodbiface.DynamoDBAPI
	tableName string
	logger    DynamoDBLogger
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

type defaultDynamoDBLogger struct{}

func (*defaultDynamoDBLogger) GetError(error, string, map[string]*dynamodb.AttributeValue) {}
func (*defaultDynamoDBLogger) GetMarshalMapError(err error, clientId string)               {}

type DynamoDBOption func(*DynamoDB)

func SetDynamoDBLogger(l DynamoDBLogger) DynamoDBOption {
	return func(d *DynamoDB) {
		if l == nil {
			d.logger = &defaultDynamoDBLogger{}
		} else {
			d.logger = l
		}
	}
}

func NewDynamoDB(
	client dynamodbiface.DynamoDBAPI,
	tableName string,
	options ...DynamoDBOption,
) *DynamoDB {
	db := &DynamoDB{
		client:    client,
		tableName: tableName,
		logger:    &defaultDynamoDBLogger{},
	}
	for _, f := range options {
		f(db)
	}

	return db
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
		d.logger.GetError(err, d.tableName, key)
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

func (d *DynamoDB) Put(ctx context.Context, product *Product) (int64, error) {
	ts := time.Now().UnixMicro()

	av, err := dynamodbattribute.MarshalMap(toDynamoDB(product, ts))
	if err != nil {
		return 0, err
	}

	if _, err := d.client.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item:      av,
	}); err != nil {
		return 0, err
	}

	return ts, nil
}
