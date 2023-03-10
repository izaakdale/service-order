package datastore

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Update(ctx context.Context, _, ticketID string) error {
	ticket, orderID, err := FetchTicket(ctx, ticketID)
	if err != nil {
		return err
	}

	ticket.QRScanned = true
	ticketMap, err := attributevalue.MarshalMap(ticket)
	if err != nil {
		return err
	}

	_, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &client.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: orderID},
			"SK": &types.AttributeValueMemberS{Value: createKey(ticketSK, ticketID)},
		},
		UpdateExpression: aws.String("set ticket = :ticket"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ticket": &types.AttributeValueMemberM{Value: ticketMap},
		},
	})
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}
	return nil
}
