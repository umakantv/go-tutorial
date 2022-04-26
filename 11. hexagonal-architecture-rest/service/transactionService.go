package service

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/dto"
	"customer_api_hex_arch/errs"
	"time"
)

type TransactionService interface {
	NewTransaction(dto.TransactionRequestDto) (*dto.TransactionResponseDto, *errs.AppError)
}

type DefaultTransactionService struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
}

func (ts DefaultTransactionService) NewTransaction(tr dto.TransactionRequestDto) (*dto.TransactionResponseDto, *errs.AppError) {
	t, e := ts.transactionRepo.NewTransaction(domain.Transaction{
		TransactionId:   "",
		AccountId:       tr.AccountId,
		Amount:          tr.Amount,
		TransactionType: tr.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	})

	if e != nil {
		return nil, e
	}

	e = ts.accountRepo.ChangeAmountWithTransaction(*t)

	if e != nil {
		return nil, e
	}

	res := t.ToTransactionResponseDto()

	return &res, nil
}

func NewTransactionService(
	transactionRepo domain.TransactionRepository,
	accountRepo domain.AccountRepository,
) DefaultTransactionService {
	return DefaultTransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}
