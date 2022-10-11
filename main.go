package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/izaakdale/service-order/service"
)

func main() {
	service := service.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_ADDRESS")), service.Router))
}
