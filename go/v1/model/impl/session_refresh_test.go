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
	if item.Role != m.Role(role) {
		t.Errorf("role must be %s but %s", role, item.Role)
	}
	if item.ExpireTime != expireTime {
		t.Errorf("expireTime must be %s but %s", expireTime, item.ExpireTime)
	}
}

func hardDeleteSessionRefresh(db *sql.DB, userId string) {
	s, _ := db.Prepare("DELETE FROM session_refresh WHERE user_id=?")
	defer s.Close()
	s.Exec(userId)
}

func TestSessionRefresh_All(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createSessionRefreshDAO(db)

	var userId string = "userId"
	var role string = "role"
	hardDeleteSessionRefresh(db, userId)
	item, err := dao.Create(userId, role)
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertSessionRefresh(t, &item, item.Token, "userId", "role", 0)

	item2, err := dao.Get(item.Token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSessionRefresh(t, &item2, item.Token, "userId", "role", item2.ExpireTime)

	userId = "NewuserId"
	role = "Newrole"
	item3, err := dao.Update(item.Token, userId, role)
	if err != nil {
		t.Errorf("Failed to Update : %s", err)
		return
	}
	assertSessionRefresh(t, &item3, item.Token, "NewuserId", "Newrole", 0)

	item4, err := dao.Get(item.Token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertSessionRefresh(t, &item4, item.Token, "NewuserId", "Newrole", item4.ExpireTime)

	item5, err := dao.Delete(item.Token)
	if err != nil {
		t.Errorf("Failed to Delete : %s", err)
		return
	}
	if !item5.IsEmpty() {
		t.Errorf("Returned item must be empty")
		return
	}
	item6, err := dao.Get(item.Token)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	if !item6.IsEmpty() {
		t.Errorf("Item must be empty")
		return
	}
}
