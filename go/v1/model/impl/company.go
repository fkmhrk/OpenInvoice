package impl

import (
	m "../"
	"database/sql"
	"errors"
	_ "fmt"
	"github.com/go-sql-driver/mysql"
)

type companyDAO struct {
	connection *Connection
}

func NewCompanyDAO(connection *Connection) *companyDAO {
	return &companyDAO{
		connection: connection,
	}
}

func (d *companyDAO) GetList() ([]*m.Company, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,name,zip,address," +
		"phone,unit FROM company " +
		"WHERE deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*m.Company
	for rows.Next() {
		list = append(list, d.readRow(rows))
	}
	return list, nil
}

func (d *companyDAO) GetById(id string) (*m.Company, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,name,zip,address," +
		"phone,unit FROM company " +
		"WHERE id=? AND deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	return d.readRow(rows), nil
}

func (d *companyDAO) Create(name, zip, address, phone, unit string) (*m.Company, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO company(" +
		"id,name,zip,address,phone,unit," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?,?,?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	// generate ID
	var id string
	for i := 0; i < 10; i++ {
		id = generateId(32)
		_, err = st.Exec(id, name, zip, address, phone, unit)
		if err == nil {
			break
		}
		id = ""
		if err2, ok := err.(*mysql.MySQLError); ok {
			if err2.Number != 1062 {
				return nil, err2
			}
		} else {
			return nil, err
		}
	}
	if len(id) == 0 {
		return nil, errors.New("Failed to create")
	}

	tr.Commit()

	return &m.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}, nil
}

func (d *companyDAO) Update(id, name, zip, address, phone, unit string) (*m.Company, error) {
	db := d.connection.Connect()

	st, err := db.Prepare("UPDATE company SET " +
		"name=?,zip=?,address=?,phone=?,unit=?," +
		"modified_time=unix_timestamp(now()) " +
		"WHERE id=? AND deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	// generate ID
	_, err = st.Exec(name, zip, address, phone, unit, id)
	if err != nil {
		return nil, err
	}

	return &m.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}, nil
}

func (d *companyDAO) readRow(rows *sql.Rows) *m.Company {
	var id, name, zip, address, phone, unit string
	rows.Scan(&id, &name, &zip, &address, &phone, &unit)
	return &m.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}
}