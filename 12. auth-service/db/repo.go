package db

import (
	"auth/logger"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	DRIVER   string
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DB       string
}

func GetDBConnection(dbConfig DatabaseConfig) *sqlx.DB {

	fmt.Println(dbConfig.DRIVER)
	// Use process env variables here instead for this
	db, err := sqlx.Open(dbConfig.DRIVER, fmt.Sprintf("%v:%v@/%v", dbConfig.USER, dbConfig.PASSWORD, dbConfig.DB))
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
