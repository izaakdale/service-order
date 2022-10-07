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

func New() *Service {
	logger.Info("Starting service...🚀")

	// TODO remove hard coding
	dao.Init("ordering-app", "eu-west-2")
	deliveryclient.Init()

	return &Service{
		Router: Routes(),
	}
}

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("POST", "/", CreateOrderHandler)
	return router
}
