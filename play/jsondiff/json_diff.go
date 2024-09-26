package json_diff

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Example struct
type Person struct {
	Name    string
	Age     int
	Email   string
	Address struct {
		City string
		Zip  int
	}
	Active bool
}

// Function to find differences and return a struct with only the differences
func diffStructs(a, b interface{}) interface{} {
	vA := reflect.ValueOf(a)
	vB := reflect.ValueOf(b)

	// Create a new instance of the struct type to store differences
	diff := reflect.New(vA.Type()).Elem()

	// Iterate over fields
	for i := 0; i < vA.NumField(); i++ {
		fieldA := vA.Field(i)
		fieldB := vB.Field(i)
		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			// Set the differing field in the diff struct
			diff.Field(i).Set(fieldB)
		}
	}

	return diff.Interface()
}

func Play() {
	person1 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Address: struct {
			City string
			Zip  int
		}{
			City: "New York",
			Zip:  10001,
		},
		Active: true,
	}
	person2 := Person{
		Name:  "Bob",
		Age:   30,
		Email: "alice@example.com",
		Address: struct {
			City string
			Zip  int
		}{
			City: "New York",
			Zip:  10002,
		},
		Active: false,
	}

	p1, _ := json.Marshal(person1)
	p2, _ := json.Marshal(person2)

	differences := diffStructs(p1, p2)

	// Marshal the differences struct to JSON
	jsonData, err := json.Marshal(differences)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
