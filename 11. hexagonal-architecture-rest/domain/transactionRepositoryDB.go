package domain

import (
	"strconv"

	"github.com/umakantv/go-utils/errs"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (tr TransactionRepositoryDB) NewTransaction(t Transaction) (*Transaction, *errs.AppError) {
	insert_query := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"

	result, err := tr.client.Exec(insert_query, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		return nil, errs.NewInternalServerError("Error while adding new transaction: " + err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, errs.NewInternalServerError("Error while getting last inserted id for transaction: " + err.Error())
	}
	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func NewTransactionRepositoryDB(client *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{
		client,
	}
}
