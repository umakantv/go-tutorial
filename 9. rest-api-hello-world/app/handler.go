package app

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

// Customer struct helps to serialize customers data to json using json.Marshal.
// Remember To Capitalize the First Letter for Keys That You Want to See in Your JSON.
// This cost me 1 hour: https://yourbasic.org/golang/gotcha-json-marshal-empty/
type Customer struct {
	Name    string    `json:"full_name"`
	City    string    `json:"city"`
	ZipCode string    `json:"zip_code"`
	t       complex64 // this will not be converted to json as it is not exported
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /hello")
	// w.Write([]byte("Hello there"))
	fmt.Fprint(w, "Hello there")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /customers")
	customers := []Customer{
		{
			"Umakant",
			"Delhi",
			"110053",
			1 + 3i,
		},
		{
			"Varun",
			"Delhi",
			"110053",
			1 + 3i,
		},
		{
			"Vikrant",
			"Delhi",
			"110053",
			1 + 3i,
		},
	}

	w.Header().Add("Content-Type", "application/json")

	// bs := make([]byte, 1024)
	// json.NewEncoder(os.Stdout).Encode(customers)
	e := json.NewEncoder(w).Encode(customers)
	if e != nil {
		// throw error here
		fmt.Fprint(w, "Server Error: ", e.Error())
	}
}

// TrimLogWriter implements io.Writer for testing the json that is getting written to the Response.
// It logs to the terminal in a trimmed fashion
type TrimLogWriter struct{}

func (TrimLogWriter) Write(p []byte) (int, error) {
	n, err := fmt.Println(
		time.Now().Format(time.UnixDate),
		string(
			p[:int(math.Min(120, float64(len(p))))],
		),
	)
	return n, err
}
