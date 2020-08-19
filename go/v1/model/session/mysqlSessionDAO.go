package session

import (
	"errors"
	"strings"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	sessionTableName = "session"
	sqlSelectAll     = "SELECT user_id,role,expire_time " +
		"FROM " + sessionTableName + " "
	sqlSelectSessionByToken = sqlSelectAll +
		"WHERE access_token=? AND deleted <> 1 " +
		"AND unix_timestamp()<expire_time LIMIT 1"
	sqlInsertSession = "INSERT INTO " + sessionTableName +
		"(access_token,user_id,role,expire_time," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,unix_timestamp()+?," +
		"unix_timestamp(),unix_timestamp(),0)"
)

type sessionDAO struct {
	connection *db.Connection
}

// NewSessionDAO creates instance
func NewSessionDAO(connection *db.Connection) model.Session {
	return &sessionDAO{
		connection: connection,
	}
}

func (d *sessionDAO) GetByToken(token string) (*entity.Session, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectSessionByToken)
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
	return &entity.Session{
		Token:      token,
		Role:       entity.Role(role),
		UserId:     userId,
		ExpireTime: expire,
	}, nil
}

func (d *sessionDAO) Create(userId, role string, expireIn int64) (*entity.Session, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsertSession)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	var token string
	for i := 0; i < 10; i++ {
		token = generateSessionID()
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
	return &entity.Session{
		Token:  token,
		UserId: userId,
		Role:   entity.Role(role),
	}, nil
}

func generateSessionID() string {
	id1 := uuid.New().String()
	id2 := uuid.New().String()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:48]
}
