package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/izaakdale/lib/listener"
	"github.com/izaakdale/lib/publisher"
	"github.com/izaakdale/service-event-order/internal/datastore"
	"github.com/izaakdale/service-event-order/pkg/proto/order"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

var (
	name = "service-event-order"
	spec specification
	gs   = &GServer{}
)

type (
	specification struct {
		QueueURL    string `envconfig:"QUEUE_URL"`
		AWSRegion   string `envconfig:"AWS_REGION"`
		AWSEndpoint string `envconfig:"AWS_ENDPOINT"`
		TopicArn    string `envconfig:"TOPIC_ARN"`
		TableName   string `envconfig:"TABLE_NAME"`
		GRPCHost    string `envconfig:"GRPC_HOST"`
		GRPCPort    string `envconfig:"GRPC_PORT"`
	}
)

func Run() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := envconfig.Process("", &spec)
	if err != nil {
		panic(err)
	}

	log.Printf("running %s\n", name)
	cfg, err := config.LoadDefaultConfig(context.Background(), func(o *config.LoadOptions) error {
		o.Region = spec.AWSRegion
		return nil
	})

	datastore.Init(getAwsDynamoClient(cfg, spec.AWSEndpoint), spec.TableName)

	srv := grpc.NewServer()
	order.RegisterOrderServiceServer(srv, gs)

	err = publisher.Initialise(cfg, spec.TopicArn, publisher.WithEndpoint(spec.AWSEndpoint))
	if err != nil {
		panic(err)
	}

	err = listener.Initialise(cfg, spec.QueueURL, listener.WithEndpoint(spec.AWSEndpoint))
	if err != nil {
		panic(err)
	}

	// want a thread stopping channel for errors in the listener
	errChan := make(chan error, 3)

	grpcServerSocket := fmt.Sprintf("%s:%s", spec.GRPCHost, spec.GRPCPort)

	// grpc server setup
	lis, err := net.Listen("tcp", grpcServerSocket)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	go func(errChan chan<- error) {
		log.Printf("listening for GRPC clients on %s\n", grpcServerSocket)
		errChan <- srv.Serve(lis)
	}(errChan)

	go listener.Listen(ctx, orderEventListener, errChan)

	// subscribe to shutdown signals
	shutCh := make(chan os.Signal, 1)
	signal.Notify(shutCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// wait on error from listener or shutdown sig.
	select {
	case err = <-errChan:
		// TODO do we definitely want to quit?
		if err != nil {
			log.Fatal(err)
		}
	case signal := <-shutCh:
		log.Printf("got shutdown signal: %s, exiting\n", signal)
		os.Exit(1)
	}
}

// allows use of localstack
func getAwsDynamoClient(cfg aws.Config, endpoint string) *dynamodb.Client {
	if endpoint != "" {
		return dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolver(dynamodb.EndpointResolverFromURL(endpoint)))
	}

	return dynamodb.NewFromConfig(cfg)
}
