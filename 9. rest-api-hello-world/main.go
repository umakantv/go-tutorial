package main

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
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	T       complex64
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:5555", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello there"))

	fmt.Println("IN HELLO")

	fmt.Fprint(w, "Hello there")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IN CUSTOMERS")
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
