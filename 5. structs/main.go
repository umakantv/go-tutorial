package main

import (
	"fmt"
)

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

func (p *person) updateFirstName(newFirstName string) {
	// (*p).firstName = newFirstName
	p.firstName = newFirstName // <-- This also works
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
}
