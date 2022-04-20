package domain

import (
	"customer_api_hex_arch/errs"
	"fmt"
	"strconv"
)

// CustomerRepositoryStub serves as the mock implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryStub struct {
	customers []Customer
}

func (crs CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {

	var customers []Customer
	for _, c := range crs.customers {
		if status != "" && c.Status == status {
			customers = append(customers, c)
		} else {
			customers = append(customers, c)
		}
	}
	fmt.Println("Customers Filtered", customers, status)
	return customers, nil
}

func (crs CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {

	for _, c := range crs.customers {
		if strconv.Itoa(c.Id) == id {
			return &c, nil
		}
	}

	return nil, errs.NewNotFoundError("Customer not found")
}

// NewCustomerRepositoryStub creates and returns a CustomerRepositoryStub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1, "Umakant", "Delhi", "110053", "1998-10-28", "active"},
		{2, "Rob", "Delhi", "110053", "1998-10-28", "inactive"},
	}

	return CustomerRepositoryStub{
		customers,
	}
}
