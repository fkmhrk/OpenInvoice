package impl

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func deleteTradingItemByTradingId(db *sql.DB, tradingId string) {
	s, _ := db.Prepare("DELETE FROM trading_item WHERE trading_id=?")
	defer s.Close()
	s.Exec(tradingId)
}

func insertTradingItem(db *sql.DB, id string, sortOrder int, tradingId, subject string, unitPrice, amount int, degree string, taxType int, memo string) {
	s, _ := db.Prepare("INSERT INTO trading_item(" +
		"id,sort_order,trading_id,subject,unit_price,amount," +
		"degree,tax_type,memo,deleted)" +
		"VALUES(?,?,?,?,?,?,?,?,?,0)")
	defer s.Close()
	s.Exec(id, sortOrder, tradingId, subject, unitPrice, amount, degree, taxType, memo)
}

func TestTradingItem0000_GetItemsById(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "user1122"
	deleteTradingItemByTradingId(db, tradingId)
	insertTradingItem(db, "item1", 1, tradingId, "subject1", 100, 2, "Yen", 1, "memo1")
	insertTradingItem(db, "item2", 2, tradingId, "subject2", 200, 4, "Yen", 1, "memo2")

	dao := createTradingDAO(db)
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Failed to get trading items by id : %s", err)
		return
	}
	if len(list) != 2 {
		t.Errorf("Wrong list length : %d", len(list))
		return
	}
	item := list[0]
	if item.Id != "item1" {
		t.Errorf("Wrong ID : %s", item.Id)
	}
	if item.SortOrder != 1 {
		t.Errorf("Wrong SortOrder : %d", item.SortOrder)
	}
	if item.Subject != "subject1" {
		t.Errorf("Wrong Subject : %s", item.Subject)
	}
	if item.UnitPrice != 100 {
		t.Errorf("Wrong UnitPrice : %d", item.UnitPrice)
	}
	if item.Amount != 2 {
		t.Errorf("Wrong Amount : %d", item.Amount)
	}
	if item.Degree != "Yen" {
		t.Errorf("Wrong Degree : %s", item.Degree)
	}
	if item.TaxType != 1 {
		t.Errorf("Wrong TaxType : %d", item.TaxType)
	}
	if item.Memo != "memo1" {
		t.Errorf("Wrong Memo : %s", item.Memo)
	}
}

func TestTradingItem0001_GetItemsById_0(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "user1122"
	deleteTradingItemByTradingId(db, tradingId)

	dao := createTradingDAO(db)
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Failed to get trading items by id : %s", err)
		return
	}
	if len(list) != 0 {
		t.Errorf("Wrong list length : %d", len(list))
		return
	}
}
