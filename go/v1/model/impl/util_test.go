package impl

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	return sql.Open("mysql", "openinvoice_test:openinvoice@/openinvoice")
}

func begin() (*sql.DB, *sql.Tx, error) {
	db, err := connect()
	if err != nil {
		return nil, nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		db.Close()
		return nil, nil, err
	}
	return db, tx, nil
}
