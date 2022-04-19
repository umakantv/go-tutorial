package app

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {

	res := TimeResponse{
		CurrentTime: time.Now().UTC().Format("2006-01-02 15:04:05 -0700 MST"),
		// CurrentTime: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
