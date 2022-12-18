package service

import (
	"customer-account-service/domain"
	"customer-account-service/dto"

	"github.com/umakantv/go-utils/errs"

	"github.com/umakantv/go-utils/logger"
)

// CustomerService defines the interface for Customer Service.
// It acts as a port for REST API adapter that communicates with our application/business/domain.
type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService implements CustomerService.
// It also adds a dependency to the repository adapter.
type DefaultCustomerService struct {
	// Note that here our service does not care whether the actual implementation is stubbed or an actual database.
	repo domain.CustomerRepository
}

// GetAllCustomers fetches and returns All customers from the repository.
func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	cs, e := s.repo.FindAll(status)

	if e != nil {
		return nil, e
	}

	var customers []dto.CustomerResponse
	for _, c := range cs {
		customers = append(customers, c.ToDto())
	}
	return customers, nil
}

// GetCustomerById fetches and returns a single customer by id from the repository.
func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, e := s.repo.ById(id)

	if e != nil {
		logger.Error(e.Error())
		return nil, e
	}

	r := c.ToDto()

	return &r, nil
}

// NewCustomerService creates and returns a CustomerService.
func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository}
}
