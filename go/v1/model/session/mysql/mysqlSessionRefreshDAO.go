package mysql

import (
	"database/sql"

	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	util "github.com/fkmhrk/OpenInvoice/v1/model/db/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
)

type session_refreshDAO struct {
	connection *db.Connection
}

func NewSessionRefreshDAO(connection *db.Connection) *session_refreshDAO {
	return &session_refreshDAO{
		connection: connection,
	}
}

func (d *session_refreshDAO) Create(userId, role string) (session.SessionRefresh, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO session_refresh(token,user_id,role,expire_time,created_time,modified_time,deleted) VALUES(?,?,?,unix_timestamp(now())+60*60*24*365,unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer st.Close()

	token, err := util.InsertWithUUID(48, func(id string) error {
		_, err = st.Exec(id, userId, role)
		return err
	})
	if err != nil {
		return session.SessionRefresh{}, err
	}

	tr.Commit()

	return session.SessionRefresh{
		Token:      token,
		UserId:     userId,
		Role:       user.Role(role),
		ExpireTime: 0,
	}, nil
}

func (d *session_refreshDAO) Get(token string) (session.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT token,user_id,role,expire_time FROM session_refresh WHERE token=? AND deleted <> 1")
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer st.Close()

	rows, err := st.Query(token)
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return session.SessionRefresh{}, nil
	}

	return d.scan(rows), nil
}

func (d *session_refreshDAO) Update(token, userId, role string) (session.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE session_refresh SET token=?, user_id=?, role=?,modified_time=unix_timestamp(now()) WHERE token=? AND deleted <> 1")
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer st.Close()

	_, err = st.Exec(token, userId, role, token)
	if err != nil {
		return session.SessionRefresh{}, err
	}

	return session.SessionRefresh{
		Token:      token,
		UserId:     userId,
		Role:       user.Role(role),
		ExpireTime: 0,
	}, nil
}

func (d *session_refreshDAO) Delete(token string) (session.SessionRefresh, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("UPDATE session_refresh SET modified_time=unix_timestamp(now()),deleted=1 WHERE token=? AND deleted <> 1")
	if err != nil {
		return session.SessionRefresh{}, err
	}
	defer st.Close()

	_, err = st.Exec(token)
	if err != nil {
		return session.SessionRefresh{}, err
	}

	return session.SessionRefresh{}, nil
}

func (d *session_refreshDAO) scan(rows *sql.Rows) session.SessionRefresh {
	var token string
	var userId string
	var role string
	var expireTime int64
	rows.Scan(&token, &userId, &role, &expireTime)
	return session.SessionRefresh{
		Token:      token,
		UserId:     userId,
		Role:       user.Role(role),
		ExpireTime: expireTime,
	}
}
