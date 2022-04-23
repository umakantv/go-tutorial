package domain

import (
	"customer_api_hex_arch/errs"
	"customer_api_hex_arch/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (r AccountRepositoryDB) NewAccount(a Account) (*Account, *errs.AppError) {
	insert_query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := r.client.Exec(insert_query, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error: " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new record: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error: " + err.Error())
	}
	a.AccountId = strconv.Itoa(int(id))
	return &a, nil
}

func NewAccountRepositoryDB(client *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client,
	}
}
