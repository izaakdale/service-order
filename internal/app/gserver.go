package app

import (
	"context"

	"github.com/izaakdale/service-event-order/internal/datastore"
	"github.com/izaakdale/service-event-order/pkg/proto/order"
)

type (
	GServer struct {
		order.OrderServiceServer
	}
)

func (gs *GServer) GetOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	o, err := datastore.FetchOrderTickets(req.OrderId)
	if err != nil {
		return nil, err
	}

	tickets := make([]*order.Ticket, 0, len(o.Tickets))
	for _, v := range o.Tickets {
		tickets = append(tickets, &order.Ticket{
			TicketId:   v.TicketID,
			FirstName:  v.FirstName,
			Surname:    v.Surname,
			QrPath:     v.QRPath,
			TicketType: v.Type,
		})
	}

	return &order.OrderResponse{
		Email:   o.Email,
		Tickets: tickets,
	}, nil
}
