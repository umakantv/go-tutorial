package main

import (
	"fmt"
)

/*
	Value Type Data Structures:
	* int
	* bool
	* string
	* float
	* struct

	Reference Type Data Structures:
	* slice
	* map
	* pointer
	* function
	* channel
*/

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	age       uint32
	contactInfo
}

func (p person) print() {
	fmt.Println(p)
}

// We need to think about pointers when we need to modify passed/original value
// Only in case Value Type data structures
func (p *person) updateFirstName(newFirstName string) {
	// (*p).firstName = newFirstName
	p.firstName = newFirstName // <-- This also works
}

// We need not think about pointers when we need to modify passed/original value
// Only in case Reference Type data structures
func updateSliceWithoutPointers(someSlice []string) {
	someSlice[1] = "who"
}

func main() {
	me := person{"Umakant", "Vashishtha", 23, contactInfo{
		"umakant@example.com", "1234567890",
	}}
	bob := person{
		firstName: "Bob",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			"bob@example.com",
			me.contactInfo.phone,
		},
	}

	bob.lastName = "Chhapri"

	var michael person
	me.print()
	bob.print()

	michael.updateFirstName("Michael")
	michael.lastName = "Scott"
	michael.age = 5
	michael.print()

	someSlice := []string{"Hey", "how", "are", "you"}

	updateSliceWithoutPointers(someSlice)
	fmt.Println(someSlice)
}
