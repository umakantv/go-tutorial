package domain

// CustomerRepositoryStub serves as the mock implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryStub struct {
	customers []Customer
}

func (crs CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return crs.customers, nil
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
