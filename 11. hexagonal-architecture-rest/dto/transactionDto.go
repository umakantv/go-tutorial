package dto

import "customer_api_hex_arch/errs"

type TransactionRequestDto struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponseDto struct {
	TransactionId string `json:"transaction_id"`
}

func (t TransactionRequestDto) Validate() *errs.AppError {
	if t.TransactionType != "deposit" && t.TransactionType != "withdrawal" {
		return errs.NewValidationError("Transaction type must be deposit or withdrawal")
	}
	return nil
}
