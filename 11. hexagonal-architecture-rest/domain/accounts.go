package domain

import (
	"customer_api_hex_arch/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerID  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      string `db:"account_balance"`
	Status      string `db:"status"`
}

type AccountRepository interface {
	NewAccount(Account) (*Account, *errs.AppError)
}
