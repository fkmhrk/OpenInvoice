package impl

import (
	m "../"
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
	st, err := db.Prepare("SELECT id,login_name,display_name," +
		"role,tel,password," +
		"created_time, modified_time FROM user " +
		"where login_name=? AND deleted <> 1 LIMIT 1")
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

	var idDB, nameDB, displayName, role, tel, passwordDB string
	var createDB, modifiedDB int64
	rows.Scan(&idDB, &nameDB, &displayName, &role, &tel,
		&passwordDB, &createDB, &modifiedDB)

	if hashPassword(password) != passwordDB {
		return nil, nil
		// return nil, errors.New("Invalid Name / Password")
	}

	return &m.User{
		Id:           idDB,
		LoginName:    nameDB,
		DisplayName:  displayName,
		Role:         m.Role(role),
		Tel:          tel,
		CreatedTime:  createDB,
		ModifiedTime: modifiedDB,
	}, nil
}

func (d *userDAO) GetList() ([]*m.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,login_name,display_name,role," +
		"created_time, modified_time FROM user " +
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

	var list []*m.User
	var id, name, display, role string
	var create, modified int64
	for rows.Next() {
		rows.Scan(&id, &name, &display, &role, &create, &modified)
		list = append(list, &m.User{
			Id:           id,
			LoginName:    name,
			DisplayName:  display,
			Role:         m.Role(role),
			CreatedTime:  create,
			ModifiedTime: modified,
		})
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
