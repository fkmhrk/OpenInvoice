package impl

import (
	"database/sql"
	"errors"
	_ "fmt"
	"strings"

	"github.com/fkmhrk/OpenInvoice/v1/model/company"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type companyDAO struct {
	connection *db.Connection
}

// New creates instance
func New(connection *db.Connection) company.DAO {
	return &companyDAO{
		connection: connection,
	}
}

func (d *companyDAO) GetList() ([]*company.Company, error) {
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

	var list []*company.Company
	for rows.Next() {
		list = append(list, d.readRow(rows))
	}
	return list, nil
}

func (d *companyDAO) GetById(id string) (*company.Company, error) {
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

func (d *companyDAO) Create(name, zip, address, phone, unit string) (*company.Company, error) {
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

	return &company.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}, nil
}

func (d *companyDAO) Update(id, name, zip, address, phone, unit string) (*company.Company, error) {
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

	return &company.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}, nil
}

func (d *companyDAO) Delete(id string) error {
	tr, err := d.connection.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()
	err = d.execDelete(tr, id)
	if err != nil {
		return err
	}
	tr.Commit()
	return nil
}

func (d *companyDAO) execDelete(tr *sql.Tx, id string) error {
	// update company id to ""
	st, err := tr.Prepare("UPDATE trading SET " +
		"company_id='',modified_time=unix_timestamp(now()) " +
		"WHERE company_id=?")
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(id)
	if err != nil {
		return err
	}

	// set delete flag
	st2, err := tr.Prepare("UPDATE company SET " +
		"deleted=1, modified_time=unix_timestamp(now()) " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer st2.Close()

	_, err = st2.Exec(id)
	return err
}

func (d *companyDAO) readRow(rows *sql.Rows) *company.Company {
	var id, name, zip, address, phone, unit string
	rows.Scan(&id, &name, &zip, &address, &phone, &unit)
	return &company.Company{
		Id:      id,
		Name:    name,
		Zip:     zip,
		Address: address,
		Phone:   phone,
		Unit:    unit,
	}
}

func generateId(length int) string {
	id1 := uuid.New().String()
	id2 := uuid.New().String()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:length]
}
