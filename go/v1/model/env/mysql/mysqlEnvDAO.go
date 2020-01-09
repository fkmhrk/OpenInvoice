package mysql

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
)

const (
	tableName     = "env"
	sqlSelectAll  = "SELECT id,value FROM " + tableName + " "
	sqlSelectByID = sqlSelectAll + "WHERE id=? AND deleted <> 1"
	sqlSelectList = sqlSelectAll + "WHERE deleted <> 1"
	sqlInsert     = "INSERT INTO " + tableName +
		"(id,value,created_time,modified_time,deleted) " +
		"VALUES(?,?,unix_timestamp(),unix_timestamp(),0)"
	sqlUpdate = "UPDATE " + tableName + " " +
		"SET value=?,modified_time=unix_timestamp() " +
		"WHERE id=? AND deleted <> 1"
	sqlSoftDelete = "UPDATE " + tableName + " " +
		"SET modified_time=unix_timestamp(now()),deleted=1 " +
		"WHERE id=? AND deleted <> 1"
)

type envDAO struct {
	connection *db.Connection
}

// New creates instance
func New(connection *db.Connection) env.DAO {
	return &envDAO{
		connection: connection,
	}
}

func (d *envDAO) Create(key, value string) (env.Env, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return env.Env{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsert)
	if err != nil {
		return env.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(key, value)
	if err != nil {
		return env.Env{}, err
	}

	tr.Commit()

	return env.Env{
		Key:   key,
		Value: value,
	}, nil
}

func (d *envDAO) Get(key string) (env.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectByID)
	if err != nil {
		return env.Env{}, err
	}
	defer st.Close()

	rows, err := st.Query(key)
	if err != nil {
		return env.Env{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return env.Env{}, nil
	}

	return d.scan(rows), nil
}

func (d *envDAO) GetList() ([]*env.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectList)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*env.Env, 0)
	for rows.Next() {
		item := d.scan(rows)
		list = append(list, &item)
	}

	return list, nil
}

func (d *envDAO) Save(list []*env.Env) error {
	tr, err := d.connection.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	createSt, err := tr.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer createSt.Close()

	updateSt, err := tr.Prepare(sqlUpdate)
	if err != nil {
		return err
	}
	defer updateSt.Close()

	for _, item := range list {
		result, err := updateSt.Exec(item.Value, item.Key)
		if err != nil {
			return err
		}
		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if count == 1 {
			continue
		}
		result, err = createSt.Exec(item.Key, item.Value)
		if err != nil {
			return err
		}
	}
	tr.Commit()
	return nil
}

func (d *envDAO) Update(key, value string) (env.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlUpdate)
	if err != nil {
		return env.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(value, key)
	if err != nil {
		return env.Env{}, err
	}

	return env.Env{
		Key:   key,
		Value: value,
	}, nil
}

func (d *envDAO) Delete(key string) (env.Env, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSoftDelete)
	if err != nil {
		return env.Env{}, err
	}
	defer st.Close()

	_, err = st.Exec(key)
	if err != nil {
		return env.Env{}, err
	}

	return env.Env{}, nil
}

func (d *envDAO) scan(rows *sql.Rows) env.Env {
	var key string
	var value string
	rows.Scan(&key, &value)
	return env.Env{
		Key:   key,
		Value: value,
	}
}
