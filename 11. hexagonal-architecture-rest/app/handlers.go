package app

import (
	"customer_api_hex_arch/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers acts as the adapter for REST client.
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, e := ch.service.GetAllCustomers()

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, e := ch.service.GetCustomerById(vars["customer_id"])

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
