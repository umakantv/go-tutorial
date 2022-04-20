package domain

import (
	"database/sql"
	"log"
	"time"
)

// CustomerRepositoryDB serves as the DB implmentation (Adapter) for customer repository (Port).
// It implements CustomerRepository interface from domain.
type CustomerRepositoryDB struct {
	db *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	query := "SELECT * FROM customers"

	rows, err := d.db.Query(query)

	if err != nil {
		log.Println("Error in fetching customers", err.Error())
	}

	var customers []Customer

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)

		if err != nil {
			log.Println("Error in scanning customers", err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, error) {

	query := "SELECT * FROM customers where customer_id = ?"

	row := d.db.QueryRow(query, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)

	if err != nil {
		log.Println("Error in scanning customers", err.Error())
		return nil, err
	}

	return &c, nil
}

// NewCustomerRepositoryDB creates and returns a CustomerRepositoryDB
func NewCustomerRepositoryDB() CustomerRepositoryDB {

	// Use process env variables here instead for this
	db, err := sql.Open("mysql", "root:12345678@/tutorial_banking")
	if err != nil {
		log.Println("Error in opening a DB connection", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error in ping to DB connection", err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDB{
		db,
	}
}
