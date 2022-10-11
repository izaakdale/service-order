package deliveryclient

import (
	"context"
	"fmt"
	"log"

	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/utils-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client *deliveryClient

type deliveryClient struct {
	delivery.DeliveryServiceClient
}

func Init(addr string) {
	logger.Info(fmt.Sprintf("Dialing %s", addr))
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
