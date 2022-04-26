package domain

import (
	"customer_api_hex_arch/errs"
	"customer_api_hex_arch/logger"
	"database/sql"
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

func (r AccountRepositoryDB) ChangeAmountWithTransaction(t Transaction) *errs.AppError {
	insert_query := "UPDATE accounts SET amount = amount + ? where account_id = ?"

	var result sql.Result
	var err error
	if t.TransactionType == "deposit" {
		result, err = r.client.Exec(insert_query, t.Amount, t.AccountId)
	} else {
		result, err = r.client.Exec(insert_query, -t.Amount, t.AccountId)
	}

	if err != nil {
		logger.Error("Error while updating account amount: " + err.Error())
		return errs.NewInternalServerError("Unexpected error: " + err.Error())
	}

	rows, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error while getting last insert id for new record: " + err.Error())
		return errs.NewInternalServerError("Unexpected error: " + err.Error())
	}

	if rows != 1 {
		logger.Error("Only one row should be affected")
		return errs.NewInternalServerError("Unexpected error: More than one row affected")
	}

	return nil
}

func NewAccountRepositoryDB(client *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client,
	}
}
