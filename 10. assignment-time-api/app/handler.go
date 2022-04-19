package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

	if tz == "" {
		tz = "Etc/UTC"
	}

	tzs := strings.Split(tz, ",")

	current_times := make(map[string]string)

	for _, tz := range tzs {
		l, e := time.LoadLocation(tz)

		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid timezone"))
			log.Println("Error in setting the timezone", e.Error())
			return
		}
		current_times[tz] = time.Now().In(l).Format("2006-01-02 15:04:05 -0700 MST")
	}

	if len(tzs) > 1 {
		log.Println("Response", current_times)
		json.NewEncoder(w).Encode(current_times)
	} else {
		res := TimeResponse{
			CurrentTime: current_times[tz],
		}

		json.NewEncoder(w).Encode(res)
	}

}
