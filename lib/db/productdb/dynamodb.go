package productdb

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

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
