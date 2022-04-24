package app

import (
	"customer_api_hex_arch/logger"
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if e := json.NewEncoder(w).Encode(data); e != nil {
		logger.Error("Unexpected error occured while encoding response: " + e.Error())
		panic(e)
	}
}
