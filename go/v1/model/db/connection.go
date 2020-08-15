package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	db *sql.DB
}

func NewConnection(db *sql.DB) *Connection {
	return &Connection{
		db: db,
	}
}

func (c *Connection) Connect() *sql.DB {
	return c.db
}

func (c *Connection) Begin() (*sql.Tx, error) {
	return c.db.Begin()
}
