package handlers

import (
	"customer_api_hex_arch/dto"
	"customer_api_hex_arch/logger"
	"customer_api_hex_arch/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers acts as the adapter for REST client.
type AccountHandlers struct {
	Service service.AccountService
}

func (ah *AccountHandlers) CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	requestDto.CustomerId = vars["customer_id"]

	logger.Info("Create Account Request", logger.Any("input", requestDto))

	e := requestDto.Validate()
	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}

	account, e := ah.Service.NewAccount(requestDto)

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
