package impl

import (
	m "../"
	"database/sql"
	"testing"
)

func createSessionRefreshDAO(db *sql.DB) *session_refreshDAO {
	return NewSessionRefreshDAO(NewConnection(db))
}

func assertSessionRefresh(t *testing.T, item *m.SessionRefresh, token, userId, role string, expireTime int64) {
	if item.Token != token {
		t.Errorf("token must be %s but %s", token, item.Token)
	}
	if item.UserId != userId {
		t.Errorf("userId must be %s but %s", userId, item.UserId)
	}
	if item.Role != role {
		t.Errorf("role must be %s but %s", role, item.Role)
	}
	if item.ExpireTime != expireTime {
		t.Errorf("expireTime must be %s but %s", expireTime, item.ExpireTime)
	}
}

func hardDeleteSessionRefresh(db *sql.DB, token string) {
	s, _ := db.Prepare("DELETE FROM session_refresh WHERE token=?")
	defer s.Close()
	s.Exec(token)
}

func TestSessionRefresh_All(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createSessionRefreshDAO(db)

	var token string = "token"
	var userId string = "userId"
	var role string = "role"
	hardDeleteSessionRefresh(db, token)
	item, err := dao.Create(token, userId, role)
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertSessionRefresh(t, &item, "token", "userId", "role", 0)

	item2, err := dao.Get(token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSessionRefresh(t, &item2, "token", "userId", "role", item2.ExpireTime)

	userId = "NewuserId"
	role = "Newrole"
	item3, err := dao.Update(token, userId, role)
	if err != nil {
		t.Errorf("Failed to Update : %s", err)
		return
	}
	assertSessionRefresh(t, &item3, "token", "NewuserId", "Newrole", 0)

	item4, err := dao.Get(token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSessionRefresh(t, &item4, "token", "NewuserId", "Newrole", item4.ExpireTime)

	item5, err := dao.Delete(token)
	if err != nil {
		t.Errorf("Failed to Delete : %s", err)
		return
	}
	if !item5.IsEmpty() {
		t.Errorf("Returned item must be empty")
		return
	}
	item6, err := dao.Get(token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	if !item6.IsEmpty() {
		t.Errorf("Item must be empty")
		return
	}
}
