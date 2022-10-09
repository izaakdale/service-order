package dao

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/izaakdale/utils-go/logger"
)

func GetOrder(pk string) (*OrderRecord, error) {

	req := struct {
		PK string `dynamodbav:"PK"`
		SK string `dynamodbav:"SK"`
	}{
		PK: OrderPrefixPK + pk,
		// SK: sk,
	}
	avs, err := dynamodbattribute.MarshalMap(&req)
	if err != nil {
		logger.Error("marshal mapping error")
		return nil, err
	}

	out, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: &client.tableName,
		Key:       avs,
	})
	if err != nil {
		logger.Error("get item error")
		return nil, err
	}

	// TODO just using status record to get things going
	var resp OrderRecord
	err = dynamodbattribute.UnmarshalMap(out.Item, &resp)
	if err != nil {
		logger.Error("unmarshal mapping error")
		return nil, err
	}

	return &resp, nil
}
