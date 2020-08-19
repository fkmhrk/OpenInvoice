package session

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	util "github.com/fkmhrk/OpenInvoice/v1/model/db/mysql"
)

const (
	refreshTableName    = "session_refresh"
	sqlSelectRefreshAll = "SELECT token,user_id,role,expire_time " +
		"FROM " + refreshTableName + " "
	sqlSelectRefreshByToken = sqlSelectRefreshAll +
		"WHERE token=? AND deleted <> 1"
	sqlInsertRefresh = "INSERT INTO " + refreshTableName +
		"(token,user_id,role,expire_time,created_time,modified_time,deleted) " +
		"VALUES(?,?,?,unix_timestamp()+60*60*24*365,unix_timestamp(),unix_timestamp(),0)"
	sqlUpdateRefresh = "UPDATE " + refreshTableName + " " +
		"SET token=?, user_id=?, role=?,modified_time=unix_timestamp() " +
		"WHERE token=? AND deleted <> 1"
	sqlSoftDeleteRefresh = "UPDATE " + refreshTableName + " " +
		"SET modified_time=unix_timestamp(),deleted=1 " +
		"WHERE token=? AND deleted <> 1"
)

type session_refreshDAO struct {
	connection *db.Connection
}

func NewSessionRefreshDAO(connection *db.Connection) model.SessionRefresh {
	return &session_refreshDAO{
		connection: connection,
	}
}

func (d *session_refreshDAO) Create(userId, role string) (entity.SessionRefresh, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsertRefresh)
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer st.Close()

	token, err := util.InsertWithUUID(48, func(id string) error {
		_, err = st.Exec(id, userId, role)
		return err
	})
	if err != nil {
		return entity.SessionRefresh{}, err
	}

	tr.Commit()

	return entity.SessionRefresh{
		Token:      token,
		UserId:     userId,
		Role:       entity.Role(role),
		ExpireTime: 0,
	}, nil
}

func (d *session_refreshDAO) Get(token string) (entity.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectRefreshByToken)
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer st.Close()

	rows, err := st.Query(token)
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.SessionRefresh{}, nil
	}

	return d.scan(rows), nil
}

func (d *session_refreshDAO) Update(token, userId, role string) (entity.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlUpdateRefresh)
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer st.Close()

	_, err = st.Exec(token, userId, role, token)
	if err != nil {
		return entity.SessionRefresh{}, err
	}

	return entity.SessionRefresh{
		Token:      token,
		UserId:     userId,
		Role:       entity.Role(role),
		ExpireTime: 0,
	}, nil
}

func (d *session_refreshDAO) Delete(token string) (entity.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSoftDeleteRefresh)
	if err != nil {
		return entity.SessionRefresh{}, err
	}
	defer st.Close()

	_, err = st.Exec(token)
	if err != nil {
		return entity.SessionRefresh{}, err
	}

	return entity.SessionRefresh{}, nil
}

func (d *session_refreshDAO) scan(rows *sql.Rows) entity.SessionRefresh {
	var token string
	var userID string
	var role string
	var expireTime int64
	rows.Scan(&token, &userID, &role, &expireTime)
	return entity.SessionRefresh{
		Token:      token,
		UserId:     userID,
		Role:       entity.Role(role),
		ExpireTime: expireTime,
	}
}
