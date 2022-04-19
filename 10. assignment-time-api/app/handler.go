package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {

	log.Println("GET /api/time")

	tz := r.URL.Query().Get("tz")
	log.Println("TZ", tz)
	w.Header().Add("Content-Type", "application/json")

	var ct string

	if tz != "" {
		l, e := time.LoadLocation(tz)

		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid timezone"))
			log.Println("Error in setting the timezone", e.Error())
			return
		}

		ct = time.Now().In(l).Format("2006-01-02 15:04:05 -0700 MST")
	} else {
		ct = time.Now().UTC().Format("2006-01-02 15:04:05 -0700 MST")
	}

	res := TimeResponse{
		CurrentTime: ct,
	}

	json.NewEncoder(w).Encode(res)
}
