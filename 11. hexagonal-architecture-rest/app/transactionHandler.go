package app

import (
	"customer_api_hex_arch/dto"
	"customer_api_hex_arch/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th *TransactionHandler) addNewTransaction(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.TransactionRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	requestDto.AccountId = vars["account_id"]

	e := requestDto.Validate()
	if e != nil {
		writeResponse(w, e.Code, e)
		return
	}
	res, e := th.service.NewTransaction(requestDto)

	if e != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}

	writeResponse(w, http.StatusCreated, res)

}
