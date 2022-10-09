package dao

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/izaakdale/utils-go/logger"
)

func FetchOrder(pk string) (*OrderRecord, error) {

	keyCond := expression.Key("PK").Equal(expression.Value(OrderPK(pk)))
	proj := expression.NamesList(
		expression.Name("SK"),
		expression.Name("items"),
		expression.Name("created"),
		expression.Name("status"),
		expression.Name("address"),
	)
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		logger.Error("expression build failure")
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(client.tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
	}

	out, err := client.QueryWithContext(context.Background(), input)
	if err != nil {
		logger.Error("query error")
		return nil, err
	}

	if *out.Count == 0 {
		logger.Error("not found error")
		return nil, ErrNoItemsFound
	}

	// fmt.Println(out.Items)

	var ret OrderRecord
	for _, item := range out.Items {
		var record OrderRecord
		if err := dynamodbattribute.UnmarshalMap(item, &record); err != nil {
			logger.Error("unmarshal error")
			return nil, err
		}
		switch record.RecordType {
		case ItemSK:
			ret.Items = record.Items
			ret.Created = record.Created
		case StatusSK:
			ret.Status = record.Status
			ret.Address = record.Address
		}
	}
	return &ret, nil
}
