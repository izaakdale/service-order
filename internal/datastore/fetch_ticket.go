package datastore

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// func FetchTicket(ctx context.Context, orderID, ticketID string) (*TicketRecord, error) {
// 	log.Printf("%+v, %+v\n", orderID, ticketID)
// 	out, err := client.GetItem(ctx, &dynamodb.GetItemInput{
// 		TableName: &client.tableName,
// 		Key: map[string]types.AttributeValue{
// 			"PK": &types.AttributeValueMemberS{Value: orderID},
// 			"SK": &types.AttributeValueMemberS{Value: ticketID},
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var ticket TicketRecord
// 	err = attributevalue.Unmarshal(out.Item[ticketSK], &ticket)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ticket, nil
// }

func FetchTicket(ctx context.Context, ticketID string) (*TicketRecord, string, error) {

	log.Printf("fetching %s\n", ticketID)
	filt := expression.Name("SK").Equal(expression.Value(ticketID))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		return nil, "", err
	}
	out, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 &client.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
	})
	if err != nil {
		return nil, "", err
	}

	var ticket TicketRecord
	err = attributevalue.Unmarshal(out.Items[0][ticketSK], &ticket)
	if err != nil {
		return nil, "", err
	}

	var orderID string
	err = attributevalue.Unmarshal(out.Items[0]["PK"], &orderID)
	if err != nil {
		return nil, "", err
	}

	return &ticket, orderID, nil
}
