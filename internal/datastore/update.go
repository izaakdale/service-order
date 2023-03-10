package datastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Update(ctx context.Context, orderID, ticketID string) error {
	ticket, err := FetchTicket(ctx, orderID, ticketID)
	if err != nil {
		return err
	}
	ticketMap, err := attributevalue.MarshalMap(ticket)
	if err != nil {
		return err
	}

	_, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &client.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: orderID},
			"SK": &types.AttributeValueMemberS{Value: ticketID},
		},
		UpdateExpression: aws.String("set ticket = :ticket"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ticket": &types.AttributeValueMemberM{Value: ticketMap},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
