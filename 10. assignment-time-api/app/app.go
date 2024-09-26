package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func getCurrentTime(w http.ResponseWriter, r *http.Request) {

	log.Println("GET", r.URL.Host, r.URL.User.Username())

	tz := r.URL.Query().Get("tz")
	log.Println("TZ", tz)
	w.Header().Add("Content-Type", "application/json")

	tzs := strings.Split(tz, ",")

	current_times := make(map[string]string, 0)

	for _, tz := range tzs {
		l, e := time.LoadLocation(tz)

		if e != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "invalid timezone %v", tz)
			// w.Write([]byte("invalid timezone"))
			log.Println("Error in setting the timezone", e.Error())
			return
		}
		ts := time.Now().In(l).Format("2006-01-02 15:04:05 -0700 MST")
		if tz == "" {
			current_times["current_time"] = ts
		} else {
			current_times[tz] = ts
		}
	}

	log.Println("Response", current_times)
	json.NewEncoder(w).Encode(current_times)

}

func Start() {
	r := mux.NewRouter()

	log.Println("Server is ready at http://localhost:5555")
	r.HandleFunc("/api/time", getCurrentTime).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:5555", r))

}
