package userlist

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestDynamoDB_Get(t *testing.T) {
	db := NewDynamoDB(&DynamoDBMock{
		GetItemFn: func(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"Email": {S: aws.String("test@example.com")},
				},
			}, nil
		},
	}, "table")

	ctx := context.Background()
	exp := "test@example.com"
	res, _, err := db.Get(ctx, exp)
	if err != nil {
		t.Fatalf("unexpected error in db.Get: %v", err)
	}

	if res.Name != exp {
		t.Errorf("Name is mismatch: exp=`%v` act=`%v`", exp, res.Name)
	}
}

func TestDynamoDB_Get_NoFile(t *testing.T) {
	db := NewDynamoDB(&DynamoDBMock{
		GetItemFn: func(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{
				Item: nil,
			}, nil
		},
	}, "table")

	ctx := context.Background()
	exp := "test@example.com"
	res, _, err := db.Get(ctx, exp)
	if err != ErrNoSuchUser {
		t.Logf("res=%v", res)
		t.Fatalf("unexpected error in db.Get: %v", err)
	}
}

func TestDynamoDB_Put(t *testing.T) {
	db := NewDynamoDB(&DynamoDBMock{
		PutItemFn: func(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			if aws.StringValue(in.TableName) != "table" {
				return nil, errors.New("table mismatch")
			}

			emailValue, ok := in.Item["Email"]
			if !ok {
				return nil, errors.New("Email field not found")
			}

			email := aws.StringValue(emailValue.S)
			if email != "test@example.com" {
				return nil, errors.New("Email is not test@example.com: value=" + email)
			}

			return &dynamodb.PutItemOutput{}, nil
		},
	}, "table")

	ctx := context.Background()
	_, err := db.Put(ctx, &User{
		Name: "test@example.com",
	})
	if err != nil {
		t.Fatalf("PutError: %v", err)
	}
}
