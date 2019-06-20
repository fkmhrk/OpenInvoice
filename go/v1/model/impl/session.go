package impl

import (
	"errors"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/go-sql-driver/mysql"
)

type sessionDAO struct {
	connection *Connection
}

func NewSessionDAO(connection *Connection) *sessionDAO {
	return &sessionDAO{
		connection: connection,
	}
}

func (d *sessionDAO) GetByToken(token string) (*m.Session, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT user_id,role,expire_time " +
		"FROM session " +
		"WHERE access_token=? AND deleted <> 1 " +
		"AND unix_timestamp(now())<expire_time LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var userId, role string
	var expire int64
	rows.Scan(&userId, &role, &expire)
	return &m.Session{
		Token:      token,
		Role:       m.Role(role),
		UserId:     userId,
		ExpireTime: expire,
	}, nil
}

func (d *sessionDAO) Create(userId, role string, expireIn int64) (*m.Session, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO session(" +
		"access_token,user_id,role,expire_time," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,unix_timestamp(now())+?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	var token string
	for i := 0; i < 10; i++ {
		token = generateSessionId()
		_, err = st.Exec(token, userId, role, expireIn)
		if err == nil {
			tr.Commit()
			break
		}
		token = ""
		if err2, ok := err.(*mysql.MySQLError); ok {
			if err2.Number != 1062 {
				return nil, err2
			}
		} else {
			return nil, err
		}
	}
	if len(token) == 0 {
		return nil, errors.New("Failed to create session")
	}
	return &m.Session{
		Token:  token,
		UserId: userId,
		Role:   m.Role(role),
	}, nil
}
