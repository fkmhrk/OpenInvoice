package impl

import (
	m "../"
	"database/sql"
	"testing"
)

func createEnvDAO(db *sql.DB) *envDAO {
	return NewEnvDAO(NewConnection(db))
}

func assertEnv(t *testing.T, item *m.Env, key, value string) {
	if item.Key != key {
		t.Errorf("key must be %s but %s", key, item.Key)
	}
	if item.Value != value {
		t.Errorf("value must be %s but %s", value, item.Value)
	}
}

func hardDeleteEnv(db *sql.DB, key string) {
	s, _ := db.Prepare("DELETE FROM env WHERE id=?")
	defer s.Close()
	s.Exec(key)
}

func TestEnv_All(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createEnvDAO(db)

	var key string = "key"
	var value string = "value"
	hardDeleteEnv(db, key)
	item, err := dao.Create(key, value)
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertEnv(t, &item, "key", "value")

	item2, err := dao.Get(key)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertEnv(t, &item2, "key", "value")

	value = "Newvalue"
	item3, err := dao.Update(key, value)
	if err != nil {
		t.Errorf("Failed to Update : %s", err)
		return
	}
	assertEnv(t, &item3, "key", "Newvalue")

	item4, err := dao.Get(key)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertEnv(t, &item4, "key", "Newvalue")

	item5, err := dao.Delete(key)
	if err != nil {
		t.Errorf("Failed to Delete : %s", err)
		return
	}
	if !item5.IsEmpty() {
		t.Errorf("Returned item must be empty")
		return
	}
	item6, err := dao.Get(key)
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	if !item6.IsEmpty() {
		t.Errorf("Item must be empty")
		return
	}
}
