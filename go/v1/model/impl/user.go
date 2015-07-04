package impl

import (
	m "../"
	"database/sql"
)

const (
	select_user = "SELECT id,login_name,display_name,role,tel,password,created_time,modified_time FROM user "
)

type userDAO struct {
	connection *Connection
}

func NewUserDAO(connection *Connection) *userDAO {
	return &userDAO{
		connection: connection,
	}
}

func (d *userDAO) GetByNamePassword(name, password string) (*m.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_user + "WHERE login_name=? AND deleted <> 1 LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	user, passwordDB, err := d.scan(rows)
	if err != nil {
		return nil, err
	}
	if hashPassword(password) != passwordDB {
		return nil, nil
		// return nil, errors.New("Invalid Name / Password")
	}

	return user, nil
}

func (d *userDAO) GetList() ([]*m.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_user + "WHERE deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*m.User, 0)
	for rows.Next() {
		item, _, err := d.scan(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (d *userDAO) Create(loginName, displayName, role, tel, password string) (*m.User, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO user(" +
		"id,login_name,display_name,role,tel,password," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?,?,?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	hashedPassword := hashPassword(password)
	id, err := insertWithUUID(32, func(id string) error {
		_, err = st.Exec(id, loginName, displayName, role, tel, hashedPassword)
		return err
	})
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &m.User{
		Id:          id,
		LoginName:   loginName,
		DisplayName: displayName,
		Role:        m.Role(role),
		Tel:         tel,
	}, nil

}

func (o *userDAO) scan(rows *sql.Rows) (*m.User, string, error) {
	var id, loginName, displayName, role, tel, password string
	var createTime, modifiedTime int64
	err := rows.Scan(&id, &loginName, &displayName, &role, &tel,
		&password, &createTime, &modifiedTime)
	if err != nil {
		return nil, "", err
	}
	return &m.User{
		Id:           id,
		LoginName:    loginName,
		DisplayName:  displayName,
		Role:         m.Role(role),
		Tel:          tel,
		CreatedTime:  createTime,
		ModifiedTime: modifiedTime,
	}, password, nil
}
