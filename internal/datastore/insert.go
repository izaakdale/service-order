package datastore

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/izaakdale/service-event/pkg/notifications"
)

func Insert(o notifications.OrderCreatedPayload) error {
	if client == nil {
		return ErrClientNotInitialised
	}
	// there will be a write request for each attendee plus one meta
	wr := make([]types.WriteRequest, 0)

	mRec := OrderRecord{
		PK: createKey(orderPK, o.OrderID),
		SK: metaSK,

		Type: metaSK,
		Metadata: MetaRecord{
			Email:       o.OrderRequest.ContactDetails.Email,
			PhoneNumber: o.OrderRequest.ContactDetails.PhoneNumber,
			Address: Address{
				NameOrNumber: o.OrderRequest.ContactDetails.Address.NameOrNumber,
				Street:       o.OrderRequest.ContactDetails.Address.Street,
				City:         o.OrderRequest.ContactDetails.Address.City,
				Postcode:     o.OrderRequest.ContactDetails.Address.Postcode,
			},
		},
	}

	recMap, err := attributevalue.MarshalMap(mRec)
	if err != nil {
		return err
	}
	wr = append(wr, types.WriteRequest{PutRequest: &types.PutRequest{Item: recMap}})

	for _, attendee := range o.OrderRequest.Attendees {
		ticketID := uuid.New().String()
		tRec := OrderRecord{
			PK: createKey(orderPK, o.OrderID),
			SK: createKey(ticketSK, ticketID),

			Type: ticketSK,
			Ticket: TicketRecord{
				TicketID:  ticketID,
				FirstName: attendee.Name,
				Surname:   attendee.Surname,
				Dob:       fmt.Sprintf("%d/%d/%d", attendee.BirthYear, attendee.BirthMonth, attendee.BirthDay),
				EventID:   o.OrderRequest.EventId,
				Type:      attendee.TicketType,
				QRPath:    generateQR(o.OrderID),
				QRScanned: false,
			},
		}

		tMap, err := attributevalue.MarshalMap(tRec)
		if err != nil {
			return err
		}
		wr = append(wr, types.WriteRequest{PutRequest: &types.PutRequest{Item: tMap}})
	}

	_, err = client.BatchWriteItem(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			client.tableName: wr,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// TODO create real QR scan path
func generateQR(id string) string {
	return fmt.Sprintf("%s/%s/%s", "http://public-address-of-ticketing-service", "qr", id)
}
