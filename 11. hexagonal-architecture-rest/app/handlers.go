package app

import (
	"customer_api_hex_arch/service"
	"encoding/json"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, e := ch.service.GetAllCustomers()

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
	}

	json.NewEncoder(w).Encode(c)
}
