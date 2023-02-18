package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/izaakdale/lib/listener"
	"github.com/izaakdale/lib/publisher"
	"github.com/kelseyhightower/envconfig"
)

var (
	name = "service-event-order"
	spec specification
)

type (
	specification struct {
		QueueURL    string `envconfig:"QUEUE_URL"`
		AWSRegion   string `envconfig:"AWS_REGION"`
		AWSEndpoint string `envconfig:"AWS_ENDPOINT"`
		TopicArn    string `envconfig:"TOPIC_ARN"`
	}
)

func Run() {
	err := envconfig.Process("", &spec)
	if err != nil {
		panic(err)
	}

	log.Printf("running %s\n", name)
	cfg, err := config.LoadDefaultConfig(context.Background(), func(o *config.LoadOptions) error {
		o.Region = spec.AWSRegion
		return nil
	})

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
	go listener.Listen(orderEventListener, errChan)

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
