package app

import (
	"customer_api_hex_arch/dto"
	"customer_api_hex_arch/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers acts as the adapter for REST client.
type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) createNewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	e := request.Validate()
	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}

	vars := mux.Vars(r)
	request.CustomerId = vars["customer_id"]

	account, e := ah.service.NewAccount(request)

	if e != nil {
		writeResponse(w, e.Code, e)
		return
	} else {
		writeResponse(w, http.StatusCreated, account)
	}
}

func NewAccountHandler(service service.AccountService) AccountHandlers {
	return AccountHandlers{service}
}
