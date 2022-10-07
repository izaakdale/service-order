package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/service-order/dao"
	deliveryclient "github.com/izaakdale/service-order/delivery_client"
	"github.com/izaakdale/service-order/pkg/model/order"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Creating order... 📝")
	var req order.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Error decoding json body")
		response.WriteJson(w, http.StatusBadRequest, response.ErrorReponse{
			Code:    http.StatusBadRequest,
			Message: "Error decoding json body",
		})
		return
	}

	orderId := uuid.New().String()
	err = dao.CreateOrder(&dao.OrderRecord{
		OrderId:    dao.QuotePrefixPK + orderId,
		RecordType: dao.ItemSK,
		Items:      req.ItemIds,
		Created:    time.Now(),
	})

	if err != nil {
		logger.Error("Error saving to the db")
		response.WriteJson(w, http.StatusInternalServerError, response.ErrorReponse{
			Code:    http.StatusInternalServerError,
			Message: "Error saving to the db",
		})
		return
	}

	// pass ID & address to delivery service
	deliveryclient.Delivery(context.Background(), &delivery.DeliveryRequest{
		OrderId: orderId,
		Address: &delivery.Address{
			HouseNameOrNumber: req.Address.HouseNameOrNumber,
			Street:            req.Address.Street,
			Town:              req.Address.Town,
			Postcode:          req.Address.Postcode,
		},
	})

	logger.Info(fmt.Sprintf("Order %s created ✅", orderId))
	response.WriteJson(w, http.StatusCreated, order.CreateOrderResponse{
		OrderId: orderId,
	})
}
