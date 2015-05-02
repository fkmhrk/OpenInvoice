package impl

import (
	m "../"
	"database/sql"
)

const (
	select_env = "SELECT id,value FROM env "
)

type envDAO struct {
	connection *Connection
}

func NewEnvDAO(connection *Connection) *envDAO {
	return &envDAO{
		connection: connection,
	}
}

func (d *envDAO) Create(key, value string) (m.Env, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return m.Env{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO env(id,value,created_time,modified_time,deleted) VALUES(?,?,unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return m.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(key, value)
	if err != nil {
		return m.Env{}, err
	}

	tr.Commit()

	return m.Env{
		Key:   key,
		Value: value,
	}, nil
}

func (d *envDAO) Get(key string) (m.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_env + "WHERE id=? AND deleted <> 1")
	if err != nil {
		return m.Env{}, err
	}
	defer st.Close()

	rows, err := st.Query(key)
	if err != nil {
		return m.Env{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return m.Env{}, nil
	}

	return d.scan(rows), nil
}

func (d *envDAO) GetList() ([]*m.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_env + "WHERE deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*m.Env, 0)
	for rows.Next() {
		item := d.scan(rows)
		list = append(list, &item)
	}

	return list, nil
}

func (d *envDAO) Update(key, value string) (m.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE env SET value=?,modified_time=unix_timestamp(now()) WHERE id=? AND deleted <> 1")
	if err != nil {
		return m.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(value, key)
	if err != nil {
		return m.Env{}, err
	}

	return m.Env{
		Key:   key,
		Value: value,
	}, nil
}

func (d *envDAO) Delete(key string) (m.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE env SET modified_time=unix_timestamp(now()),deleted=1 WHERE id=? AND deleted <> 1")
	if err != nil {
		return m.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(key)
	if err != nil {
		return m.Env{}, err
	}

	return m.Env{}, nil
}

func (d *envDAO) scan(rows *sql.Rows) m.Env {
	var key string
	var value string
	rows.Scan(&key, &value)
	return m.Env{
		Key:   key,
		Value: value,
	}
}
