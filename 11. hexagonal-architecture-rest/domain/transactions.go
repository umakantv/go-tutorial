package domain

import (
	"customer-account-service/dto"

	"github.com/umakantv/go-utils/errs"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) ToTransactionResponseDto() dto.TransactionResponseDto {
	return dto.TransactionResponseDto{
		TransactionId: t.TransactionId,
	}
}

type TransactionRepository interface {
	NewTransaction(Transaction) (*Transaction, *errs.AppError)
}
