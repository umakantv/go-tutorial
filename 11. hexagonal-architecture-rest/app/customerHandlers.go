package app

import (
	"customer_api_hex_arch/service"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers acts as the adapter for REST client.
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	c, e := ch.service.GetAllCustomers(status)

	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}
	writeResponse(w, http.StatusOK, c)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, e := ch.service.GetCustomerById(vars["customer_id"])

	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}
	writeResponse(w, http.StatusOK, c)
}
