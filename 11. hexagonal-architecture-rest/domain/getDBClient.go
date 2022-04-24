package domain

import (
	"customer_api_hex_arch/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetDBConnection() *sqlx.DB {

	// Use process env variables here instead for this
	db, err := sqlx.Open("mysql", "root:12345678@/tutorial_banking")
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

	return db
}
