package dao

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func PutOrder(o *OrderRecord) error {

	orderMap, err := dynamodbattribute.MarshalMap(o)
	if err != nil {
		return err
	}

	writeRequest := []*dynamodb.WriteRequest{}
	writeRequest = append(writeRequest, &dynamodb.WriteRequest{
		PutRequest: &dynamodb.PutRequest{
			Item: orderMap,
		},
	})

	_, err = client.BatchWriteItemWithContext(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			client.tableName: writeRequest,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
