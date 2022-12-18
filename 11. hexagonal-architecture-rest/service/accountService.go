package service

import (
	"customer-account-service/domain"
	"customer-account-service/dto"
	"time"

	"github.com/umakantv/go-utils/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepositoryDB
}

func (s DefaultAccountService) NewAccount(a dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	newAccount, err := s.repo.NewAccount(domain.Account{
		AccountId:   "", // will be set later
		CustomerID:  a.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	})

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepositoryDB) AccountService {
	return DefaultAccountService{
		repo,
	}
}
