package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	testdb "github.com/fkmhrk/OpenInvoice/v1/model/db/test"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	_ "github.com/go-sql-driver/mysql"
)

func createSessionDAO(sqlDB *sql.DB) session.SessionDAO {
	c := db.NewConnection(sqlDB)
	return NewSessionDAO(c)
}

func deleteByToken(db *sql.DB, token string) {
	s, _ := db.Prepare("DELETE FROM session WHERE access_token=?")
	defer s.Close()
	s.Exec(token)
}

func deleteSessionByUserId(db *sql.DB, id string) {
	s, _ := db.Prepare("DELETE FROM session WHERE user_id=?")
	defer s.Close()
	s.Exec(id)
}

func insertSession(db *sql.DB, token, id string, expire int64) {
	s, err := db.Prepare("INSERT INTO session" +
		"(access_token,user_id,expire_time,deleted)" +
		"VALUES(?,?,unix_timestamp(now())+?,0)")
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer s.Close()
	s.Exec(token, id, expire)
}

func TestSessionDAO_0000_GetByToken(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	token := "token1234"
	id := "testID0000"
	var expire int64 = 1000
	deleteByToken(db, token)
	insertSession(db, token, id, expire)

	dao := createSessionDAO(db)
	session, err := dao.GetByToken(token)
	if err != nil {
		t.Errorf("Failed to get session : %s", err)
		return
	}
	if session.UserId != id {
		t.Errorf("Invalid User ID %s", session.UserId)
		return
	}
}

func TestSessionDAO_0100_Create_Get(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	id := "testID0000"
	role := "Read"
	var expireIn int64 = 60 * 1000
	deleteSessionByUserId(db, id)

	dao := createSessionDAO(db)
	session, err := dao.Create(id, role, expireIn)
	if err != nil {
		t.Errorf("Failed to create session : %s", err)
		return
	}
	if session.UserId != id {
		t.Errorf("Invalid User ID %s", session.UserId)
		return
	}
	if string(session.Role) != role {
		t.Errorf("Role must be %s but %s", role, string(session.Role))
		return
	}

	// get
	session, err = dao.GetByToken(session.Token)
	if err != nil {
		t.Errorf("Failed to get session : %s", err)
		return
	}
	if session.UserId != id {
		t.Errorf("Invalid User ID %s", session.UserId)
		return
	}
	if string(session.Role) != role {
		t.Errorf("Role must be %s but %s", role, string(session.Role))
		return
	}
}
