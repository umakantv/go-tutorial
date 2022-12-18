package domain

import (
	"customer-account-service/dto"

	"github.com/umakantv/go-utils/errs"
)

// Customer type represents the customer data.
type Customer struct {
	Id          int `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string // whether active or inactive
}

// statusAsText returns status as text
func (c Customer) statusAsText() string {
	s := "active"
	if c.Status == "0" {
		s = "inactive"
	}
	return s
}

// ToDto returns Data Transfer Object for API
func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository is an interface for implementing store dependencies.
// This acts as a Secondary Port for any adapter services that are interested in behaving as a store for customers.
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
