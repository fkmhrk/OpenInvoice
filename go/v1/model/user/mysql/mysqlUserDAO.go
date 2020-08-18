package mysql

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	util "github.com/fkmhrk/OpenInvoice/v1/model/db/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"
)

const (
	tableName    = "user"
	sqlSelectAll = "SELECT id,login_name,display_name,role,tel,password,created_time,modified_time " +
		"FROM " + tableName + " "
	sqlSelectByNamePassword = sqlSelectAll +
		"WHERE login_name=? AND deleted <> 1 LIMIT 1"
	sqlSelectList = sqlSelectAll +
		"WHERE deleted <> 1"
	sqlSelectByID = sqlSelectAll + "WHERE id=? AND deleted <> 1"
	sqlInsert     = "INSERT INTO " + tableName +
		"(id,login_name,display_name,role,tel,password," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?,?,?," +
		"unix_timestamp(),unix_timestamp(),0)"
	sqlUpdateWithoutPassword = "UPDATE " + tableName + " " +
		"SET login_name=?,display_name=?,tel=?," +
		"modified_time=unix_timestamp() " +
		"WHERE id=? AND deleted<>1"
	sqlUpdateWithPassword = "UPDATE " + tableName + " " +
		"SET login_name=?,display_name=?,tel=?,password=?," +
		"modified_time=unix_timestamp() " +
		"WHERE id=? AND deleted<>1"
	sqlUpdateTradingAssignee = "UPDATE trading " +
		"SET assignee='empty',modified_time=unix_timestamp() " +
		"WHERE assignee=?"
	sqlSoftDelete = "UPDATE " + tableName + " " +
		"SET deleted=1,modified_time=unix_timestamp() " +
		"WHERE id=?"
)

type userDAO struct {
	connection *db.Connection
}

func New(connection *db.Connection) model.User {
	return &userDAO{
		connection: connection,
	}
}

func (d *userDAO) GetByNamePassword(name, password string) (*entity.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectByNamePassword)
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
	}

	return user, nil
}

func (d *userDAO) GetList() ([]*entity.User, error) {
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

	list := make([]*entity.User, 0)
	for rows.Next() {
		item, _, err := d.scan(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (d *userDAO) GetById(id string) (*entity.User, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectByID)
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

	user, _, err := d.scan(rows)
	return user, err
}

func (d *userDAO) Create(loginName, displayName, role, tel, password string) (*entity.User, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsert)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	hashedPassword := hashPassword(password)
	id, err := util.InsertWithUUID(32, func(id string) error {
		_, err = st.Exec(id, loginName, displayName, role, tel, hashedPassword)
		return err
	})
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &entity.User{
		Id:          id,
		LoginName:   loginName,
		DisplayName: displayName,
		Role:        entity.Role(role),
		Tel:         tel,
	}, nil

}

func (o *userDAO) Update(id, loginName, displayName, role, tel, password string) (*entity.User, error) {
	if len(password) == 0 {
		return o.updateWithoutPassword(id, loginName, displayName, role, tel)
	} else {
		return o.updateWithPassword(id, loginName, displayName, role, tel, password)
	}
}

func (o *userDAO) Delete(id string) error {
	tr, err := o.connection.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	err = o.updateTradingAssignee(tr, id)
	if err != nil {
		return err
	}

	err = o.deleteUser(tr, id)
	if err != nil {
		return err
	}

	tr.Commit()
	return nil
}

func (o *userDAO) updateWithoutPassword(id, loginName, displayName, role, tel string) (*entity.User, error) {
	tr, err := o.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlUpdateWithoutPassword)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(loginName, displayName, tel, id)
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &entity.User{
		Id:          id,
		LoginName:   loginName,
		DisplayName: displayName,
		Role:        entity.Role(role),
		Tel:         tel,
	}, nil
}

func (o *userDAO) updateWithPassword(id, loginName, displayName, role, tel, password string) (*entity.User, error) {
	tr, err := o.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlUpdateWithPassword)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	hashedPassword := hashPassword(password)
	_, err = st.Exec(loginName, displayName, tel, hashedPassword, id)
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &entity.User{
		Id:          id,
		LoginName:   loginName,
		DisplayName: displayName,
		Role:        entity.Role(role),
		Tel:         tel,
	}, nil
}

func (o *userDAO) updateTradingAssignee(tr *sql.Tx, id string) error {
	st, err := tr.Prepare(sqlUpdateTradingAssignee)
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(id)
	return err
}

func (o *userDAO) deleteUser(tr *sql.Tx, id string) error {
	st, err := tr.Prepare(sqlSoftDelete)
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(id)
	return err
}

func (o *userDAO) scan(rows *sql.Rows) (*entity.User, string, error) {
	var id, loginName, displayName, role, tel, password string
	var createTime, modifiedTime int64
	err := rows.Scan(&id, &loginName, &displayName, &role, &tel,
		&password, &createTime, &modifiedTime)
	if err != nil {
		return nil, "", err
	}
	return &entity.User{
		Id:           id,
		LoginName:    loginName,
		DisplayName:  displayName,
		Role:         entity.Role(role),
		Tel:          tel,
		CreatedTime:  createTime,
		ModifiedTime: modifiedTime,
	}, password, nil
}
