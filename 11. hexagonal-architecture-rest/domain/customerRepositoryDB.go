package domain

import (
	"customer_api_hex_arch/errs"
	"customer_api_hex_arch/logger"
	"database/sql"
	"time"
)

// CustomerRepositoryDB serves as the DB implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryDB struct {
	db *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	query := "SELECT * FROM customers"

	if status == "active" {
		query = "SELECT * FROM customers where status = 1"
	} else if status == "inactive" {
		query = "SELECT * FROM customers where status = 0"
	}

	// fmt.Println(query)

	rows, err := d.db.Query(query)

	if err != nil {
		logger.Error("Error in fetching customers " + err.Error())
	}

	var customers []Customer

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)

		if err != nil {
			logger.Error("Error in scanning customers " + err.Error())
			return nil, errs.NewInternalServerError(err.Error())
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {

	query := "SELECT * FROM customers where customer_id = ?"

	row := d.db.QueryRow(query, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)

	if err != nil {

		logger.Info("Error in scanning customers for id: " + id + " - " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &c, nil
}

// NewCustomerRepositoryDB creates and returns a CustomerRepositoryDB
func NewCustomerRepositoryDB() CustomerRepositoryDB {

	// Use process env variables here instead for this
	db, err := sql.Open("mysql", "root:12345678@/tutorial_banking")
	if err != nil {
		logger.Error("Error in opening a DB connection " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		logger.Error("Error in ping to DB connection " + err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDB{
		db,
	}
}
