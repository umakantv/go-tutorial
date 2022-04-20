package domain

import "customer_api_hex_arch/errs"

// Customer type represents the customer data.
type Customer struct {
	Id          int `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string // whether active or inactive
}

// CustomerRepository is an interface for implementing store dependencies.
// This acts as a Secondary Port for any adapter services that are interested in behaving as a store for customers.
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
