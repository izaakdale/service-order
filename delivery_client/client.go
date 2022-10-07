package deliveryclient

import (
	"context"
	"log"

	"github.com/izaakdale/service-delivery/model/delivery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client *deliveryClient

type deliveryClient struct {
	delivery.DeliveryServiceClient
}

func Init() {

	// TODO remove hard coding
	addr := "localhost:50001"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s", addr)
	}
	c := delivery.NewDeliveryServiceClient(conn)

	client = &deliveryClient{
		DeliveryServiceClient: c,
	}
}

func Delivery(ctx context.Context, in *delivery.DeliveryRequest) {
	client.DeliveryServiceClient.Delivery(ctx, in)
}
