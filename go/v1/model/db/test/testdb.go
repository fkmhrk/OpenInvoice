package testdb

import (
	"database/sql"
	// Use for sql.Connect()
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/openinvoice")
}
