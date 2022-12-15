package handlers

import (
	"customer_api_hex_arch/logger"
	"customer_api_hex_arch/service"

	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers acts as the adapter for REST client.
type CustomerHandlers struct {
	Service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	logger.Info("Get All Customers")

	status := r.URL.Query().Get("status")
	c, e := ch.Service.GetAllCustomers(status)

	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}
	writeResponse(w, http.StatusOK, c)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, e := ch.Service.GetCustomerById(vars["customer_id"])

	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}
	writeResponse(w, http.StatusOK, c)
}
