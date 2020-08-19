package env

import (
	"database/sql"
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	testdb "github.com/fkmhrk/OpenInvoice/v1/model/db/test"
)

func createEnvDAO(sqlDB *sql.DB) model.Env {
	return New(db.NewConnection(sqlDB))
}

func assertEnv(t *testing.T, item *entity.Env, key, value string) {
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
	db, err := testdb.Connect()
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

func TestEnv_GetList(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createEnvDAO(db)

	hardDeleteEnv(db, "key1")
	hardDeleteEnv(db, "key2")
	hardDeleteEnv(db, "key3")

	item, err := dao.Create("key1", "value1")
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertEnv(t, &item, "key1", "value1")

	item, err = dao.Create("key2", "value2")
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertEnv(t, &item, "key2", "value2")

	// get List
	list1, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to Get List : %s", err)
		return
	}

	// Add
	item, err = dao.Create("key3", "value3")
	if err != nil {
		t.Errorf("Failed to Create : %s", err)
		return
	}
	assertEnv(t, &item, "key3", "value3")

	list2, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to Get List : %s", err)
		return
	}
	if len(list2) != len(list1)+1 {
		t.Errorf("Unexpected length list1=%d list2=%d", len(list1), len(list2))
	}

}

func TestEnv_Save(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()

	dao := createEnvDAO(db)

	hardDeleteEnv(db, "key1")
	hardDeleteEnv(db, "key2")
	hardDeleteEnv(db, "key3")

	list := []*entity.Env{
		{
			Key:   "key1",
			Value: "value1",
		},
		{
			Key:   "key2",
			Value: "value2",
		},
	}

	err = dao.Save(list)
	if err != nil {
		t.Errorf("Failed to Save : %s", err)
		return
	}

	// get List
	list2, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to Get List : %s", err)
		return
	}
	if len(list2) != 2 {
		t.Errorf("List must be 2 but %d", len(list2))
		return
	}

	list = []*entity.Env{
		{
			Key:   "key1",
			Value: "value3",
		},
		{
			Key:   "key3",
			Value: "value4",
		},
	}

	err = dao.Save(list)
	if err != nil {
		t.Errorf("Failed to Save : %s", err)
		return
	}

	// check
	item, err := dao.Get("key1")
	if err != nil {
		t.Errorf("Failed to Get : %s", err)
		return
	}
	assertEnv(t, &item, "key1", "value3")

	list3, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to Get List : %s", err)
		return
	}
	if len(list3) != 3 {
		t.Errorf("List must be 3 but %d", len(list3))
		return
	}
}
