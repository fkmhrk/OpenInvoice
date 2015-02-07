package impl

import (
	m "../"
	_ "errors"
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
	st, err := db.Prepare("SELECT id,login_name,password," +
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

	var idDB, nameDB, passwordDB string
	var createDB, modifiedDB int64
	rows.Scan(&idDB, &nameDB, &passwordDB, &createDB, &modifiedDB)

	if hashPassword(password) != passwordDB {
		return nil, nil
		// return nil, errors.New("Invalid Name / Password")
	}

	return &m.User{
		Id:           idDB,
		LoginName:    nameDB,
		CreatedTime:  createDB,
		ModifiedTime: modifiedDB,
	}, nil
}

func (d *userDAO) GetList() ([]*m.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,login_name,display_name," +
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
	var id, name, display string
	var create, modified int64
	for rows.Next() {
		rows.Scan(&id, &name, &display, &create, &modified)
		list = append(list, &m.User{
			Id:           id,
			LoginName:    name,
			DisplayName:  display,
			CreatedTime:  create,
			ModifiedTime: modified,
		})
	}
	return list, nil
}
