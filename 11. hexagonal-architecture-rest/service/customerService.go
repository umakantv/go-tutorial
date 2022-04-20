package service

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/errs"
)

// CustomerService defines the interface for Customer Service.
// It acts as a port for REST API adapter that communicates with our application/business/domain.
type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService implements CustomerService.
// It also adds a dependency to the repository adapter.
type DefaultCustomerService struct {
	// Note that here our service does not care whether the actual implementation is stubbed or an actual database.
	repo domain.CustomerRepository
}

// GetAllCustomers fetches and returns All customers from the repository.
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}

// GetCustomerById fetches and returns a single customer by id from the repository.
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// NewCustomerService creates and returns a CustomerService.
func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository}
}
