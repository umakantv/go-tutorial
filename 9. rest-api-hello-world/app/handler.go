package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Customer struct helps to serialize customers data to json using json.Marshal.
// Remember To Capitalize the First Letter for Keys That You Want to See in Your JSON.
// This cost me 1 hour: https://yourbasic.org/golang/gotcha-json-marshal-empty/
type Customer struct {
	Id      int       `json:"id"`
	Name    string    `json:"full_name"`
	City    string    `json:"city"`
	ZipCode string    `json:"zip_code"`
	t       complex64 // this will not be converted to json as it is not exported
}

var customers = []Customer{
	{
		1,
		"Umakant",
		"Delhi",
		"110053",
		1 + 3i,
	},
	{
		2,
		"Varun",
		"Delhi",
		"110053",
		1 + 3i,
	},
	{
		3,
		"Vikrant",
		"Delhi",
		"110053",
		1 + 3i,
	},
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /hello")
	// w.Write([]byte("Hello there"))
	fmt.Fprint(w, "Hello there")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /customers")
	writeJsonOrXml(w, r, customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /customers")

	id, e := strconv.Atoi((mux.Vars(r)["customer_id"]))
	if e != nil {
		log.Panicln("Bad input data for customer id.", e.Error())
		fmt.Fprint(w, "Bad input data for customer id.")
		return
	}
	for _, customer := range customers {
		if customer.Id == id {
			writeJsonOrXml(w, r, customer)
			return
		}
	}
	fmt.Fprint(w, "Not Found")
}

func writeJsonOrXml(w http.ResponseWriter, r *http.Request, data any) {
	var e error

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		e = xml.NewEncoder(w).Encode(data)
	} else {
		w.Header().Add("Content-Type", "application/json")
		e = json.NewEncoder(w).Encode(data)
	}

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
