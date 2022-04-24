package dto

type TransactionRequestDto struct {
	AccountId       int     `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponseDto struct {
	TransactionId string `json:"transaction_id"`
}
