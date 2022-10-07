package main

import (
	"log"
	"net/http"

	"github.com/izaakdale/service-order/service"
)

func main() {
	service := service.New()
	log.Fatal(http.ListenAndServe("localhost:8080", service.Router))
}
