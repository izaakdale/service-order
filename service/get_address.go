package service

import (
	"encoding/json"
	"net/http"

	"github.com/izaakdale/service-order/dao"
	"github.com/izaakdale/service-order/pkg/model/order"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

func getAddressHandler(w http.ResponseWriter, r *http.Request) {
	var req order.GetOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Error decoding json body")
		response.WriteJson(w, http.StatusBadRequest, response.NewBadRequestError(err.Error()))
		return
	}

	o, err := dao.FetchAddress(req.OrderId)
	if err != nil {
		logger.Error("Error getting from the db")
		if err == dao.ErrNoItemsFound {
			response.WriteJson(w, http.StatusNotFound, response.NewNotFoundError(err.Error()))
			return
		}
		response.WriteJson(w, http.StatusInternalServerError, response.NewInternalError(err.Error()))
		return
	}

	response.WriteJson(w, http.StatusOK, o)

}
