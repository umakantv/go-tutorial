package domain

import (
	"customer-account-service/dto"

	"github.com/umakantv/go-utils/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"account_balance"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	NewAccount(Account) (*Account, *errs.AppError)
	ChangeAmountWithTransaction(Transaction) *errs.AppError
}
