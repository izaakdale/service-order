package datastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func FetchTicket(ctx context.Context, orderID, ticketID string) (*TicketRecord, error) {
	out, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &client.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: orderID},
			"SK": &types.AttributeValueMemberS{Value: ticketID},
		},
	})
	if err != nil {
		return nil, err
	}

	var ticket TicketRecord
	err = attributevalue.Unmarshal(out.Item, &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}
