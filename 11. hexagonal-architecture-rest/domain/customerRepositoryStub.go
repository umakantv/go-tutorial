package domain

import (
	"errors"
	"strconv"
)

// CustomerRepositoryStub serves as the mock implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryStub struct {
	customers []Customer
}

func (crs CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return crs.customers, nil
}

func (crs CustomerRepositoryStub) ById(id string) (*Customer, error) {

	for _, c := range crs.customers {
		if strconv.Itoa(c.Id) == id {
			return &c, nil
		}
	}

	return nil, errors.New("Customer not found")
}

// NewCustomerRepositoryStub creates and returns a CustomerRepositoryStub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1, "Umakant", "Delhi", "110053", "1998-10-28", "active"},
		{2, "Rob", "Delhi", "110053", "1998-10-28", "active"},
	}

	return CustomerRepositoryStub{
		customers,
	}
}
