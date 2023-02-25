package datastore

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func FetchOrderTickets(id string) (*OrderTickets, error) {
	if client == nil {
		return nil, ErrClientNotInitialised
	}

	log.Printf("fetching item %s\n", id)

	keyCond := expression.Key("PK").Equal(expression.Value(createKey(orderPK, id)))
	proj := expression.NamesList(
		expression.Name("metadata"),
		expression.Name("ticket"),
		expression.Name("type"),
	)
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 &client.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
	}

	out, err := client.Query(context.Background(), input)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errors.New("empty dynamo response")
	}

	resp := OrderTickets{}

	for _, item := range out.Items {
		var rec OrderRecord
		if err := attributevalue.UnmarshalMap(item, &rec); err != nil {
			return nil, err
		}

		switch rec.Type {
		case metaSK:
			metaRec := MetaRecord{}
			err := attributevalue.Unmarshal(item[metaSK], &metaRec)
			if err != nil {
				return nil, err
			}
			resp.Email = metaRec.Email
		case ticketSK:
			ticketRec := TicketRecord{}
			err := attributevalue.Unmarshal(item[ticketSK], &ticketRec)
			if err != nil {
				return nil, err
			}
			resp.Tickets = append(resp.Tickets, &ticketRec)
		default:
			log.Printf("hit default switch while fetching %s\n", id)
		}
	}

	return &resp, nil
}
