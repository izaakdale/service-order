package service

import (
	"github.com/izaakdale/service-order/dao"
	deliveryclient "github.com/izaakdale/service-order/delivery_client"
	"github.com/izaakdale/utils-go/logger"
	"github.com/julienschmidt/httprouter"
)

type Service struct {
	Router *httprouter.Router
}

func New(tableName, region, deliverSrvAddr string) *Service {

	logger.Info("Starting service...🚀")

	dao.Init(tableName, region)
	deliveryclient.Init(deliverSrvAddr)

	return &Service{
		Router: Routes(),
	}
}

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("POST", "/", createOrderHandler)
	router.HandlerFunc("GET", "/", getOrderHandler)
	router.HandlerFunc("GET", "/status", getStatusHandler)
	router.HandlerFunc("GET", "/items", getItemsHandler)
	router.HandlerFunc("GET", "/address", getAddressHandler)
	return router
}
