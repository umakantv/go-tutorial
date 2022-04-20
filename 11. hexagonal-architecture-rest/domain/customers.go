package domain

// Customer type represents the customer data.
type Customer struct {
	Id          int
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string // whether active or inactive
}

// CustomerRepository is an interface for implementing store dependencies.
// This acts as a Secondary Port for any adapter services that are interested in behaving as a store for customers.
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}
