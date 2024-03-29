package dto

import (
	"log"

	"github.com/umakantv/go-utils/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id"`
}

func (r NewAccountResponse) Debug() {
	log.Println("AccountID", r.AccountId)
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("initial amount must be atleast 5000.00")
	} else if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("account type must be saving or checking")
	}
	return nil
}
