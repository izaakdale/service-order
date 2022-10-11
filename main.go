package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/izaakdale/service-order/service"
	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	Address         string
	DynamoTableName string
	DynamoRegion    string
	GrpcAddress     string
}

func main() {

	var s Spec
	err := envconfig.Process("service", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	service := service.New(s.DynamoTableName, s.DynamoRegion, s.GrpcAddress)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Address), service.Router))
}
