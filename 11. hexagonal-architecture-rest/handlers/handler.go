package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/umakantv/go-utils/logger"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if e := json.NewEncoder(w).Encode(data); e != nil {
		logger.Error("Unexpected error occured while encoding response: " + e.Error())
		panic(e)
	}
}
