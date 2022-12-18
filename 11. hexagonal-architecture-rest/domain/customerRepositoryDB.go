package domain

import (
	"database/sql"

	"github.com/umakantv/go-utils/errs"

	"github.com/umakantv/go-utils/logger"

	"github.com/jmoiron/sqlx"
)

// CustomerRepositoryDB serves as the DB implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	query := "SELECT * FROM customers"

	if status == "active" {
		query = "SELECT * FROM customers where status = 1"
	} else if status == "inactive" {
		query = "SELECT * FROM customers where status = 0"
	}

	var customers []Customer
	err := d.db.Select(&customers, query)

	if err != nil {
		logger.Error("Error in fetching customers " + err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {

	query := "SELECT * FROM customers where customer_id = ?"

	var c Customer
	err := d.db.Get(&c, query, id)

	if err != nil {
		logger.Info("Error in scanning customers")
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		return nil, errs.NewInternalServerError("Unexpected error: " + err.Error())
	}

	return &c, nil
}

// NewCustomerRepositoryDB creates and returns a CustomerRepositoryDB
func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{
		db,
	}
}
