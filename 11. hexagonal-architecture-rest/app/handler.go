package app

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if e := json.NewEncoder(w).Encode(data); e != nil {
		panic(e)
	}
}
