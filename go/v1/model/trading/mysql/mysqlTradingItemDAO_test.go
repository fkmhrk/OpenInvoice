package mysql

import (
	"database/sql"
	"testing"

	testdb "github.com/fkmhrk/OpenInvoice/v1/model/db/test"
	"github.com/fkmhrk/OpenInvoice/v1/model/trading"
	_ "github.com/go-sql-driver/mysql"
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
	db, err := testdb.Connect()
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
		t.Errorf("Wrong Amount : %f", item.Amount)
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
	db, err := testdb.Connect()
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

func TestTradingItem0100_CreateItem(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "trade1122"
	deleteTradingItemByTradingId(db, tradingId)

	dao := createTradingDAO(db)
	item, err := dao.CreateItem(tradingId, "Subject1", "M/D", "Memo1",
		1, 100, 5, 1)
	if err != nil {
		t.Errorf("Unexpected error : %s", err)
		return
	}
	if item.TradingId != "trade1122" {
		t.Errorf("Unexpected trading ID : %s", item.TradingId)
	}

	// check
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Unexpected getList error : %s", err)
		return
	}
	if len(list) != 1 {
		t.Errorf("Unexpected list length : %d", len(list))
		return
	}
	assertTradingItem(t, list[0], item.Id, "trade1122", "Subject1", 1,
		100, 5, "M/D", 1, "Memo1")
}

func TestTradingItem0101_CreateItem_2(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "trade1122"
	deleteTradingItemByTradingId(db, tradingId)

	dao := createTradingDAO(db)
	for i := 0; i < 2; i++ {
		_, err := dao.CreateItem(tradingId, "Subject1", "M/D", "Memo1",
			i, 100*i, float64(5*i), 1)
		if err != nil {
			t.Errorf("Unexpected error : %s", err)
			return
		}
	}

	// check
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Unexpected getList error : %s", err)
		return
	}
	if len(list) != 2 {
		t.Errorf("Unexpected list length : %d", len(list))
		return
	}
}
func TestTradingItem0200_UpdateItem(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "trade1122"
	deleteTradingItemByTradingId(db, tradingId)

	dao := createTradingDAO(db)
	item, err := dao.CreateItem(tradingId, "Subject1", "M/D", "Memo1",
		1, 100, 5, 1)
	if err != nil {
		t.Errorf("Unexpected error : %s", err)
		return
	}

	// update
	item2, err := dao.UpdateItem(item.Id, tradingId, "Subject2", "M/M",
		"NewMemo", 2, 200, 10, 2)
	if err != nil {
		t.Errorf("Update is failed : %s", err)
		return
	}

	// check
	assertTradingItem(t, item2, item.Id, "trade1122", "Subject2", 2,
		200, 10, "M/M", 2, "NewMemo")

	// re-check
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Unexpected getList error : %s", err)
		return
	}
	if len(list) != 1 {
		t.Errorf("Unexpected list length : %d", len(list))
		return
	}

	assertTradingItem(t, item2, item.Id, "trade1122", "Subject2", 2,
		200, 10, "M/M", 2, "NewMemo")
}

func TestTradingItem0300_DeleteItem(t *testing.T) {
	db, err := testdb.Connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	tradingId := "trade1122"
	deleteTradingItemByTradingId(db, tradingId)

	dao := createTradingDAO(db)
	item, err := dao.CreateItem(tradingId, "Subject1", "M/D", "Memo1",
		1, 100, 5, 1)
	if err != nil {
		t.Errorf("Unexpected error : %s", err)
		return
	}

	// delete
	err = dao.SoftDeleteItem(item.Id, tradingId)
	if err != nil {
		t.Errorf("DeleteItem is failed : %s", err)
		return
	}

	// check
	list, err := dao.GetItemsById(tradingId)
	if err != nil {
		t.Errorf("Unexpected getList error : %s", err)
		return
	}
	if len(list) != 0 {
		t.Errorf("Unexpected list length : %d", len(list))
		return
	}
}
func assertTradingItem(t *testing.T, item *trading.TradingItem,
	id, tradingId, subject string, sortOrder, unitPrice int, amount float64,
	degree string, taxType int, memo string) {
	caller := getCaller()
	if item.Id != id {
		t.Errorf("%s Id must be %s but %s", caller, id, item.Id)
	}
	if item.TradingId != tradingId {
		t.Errorf("%s TradingId must be %s but %s", caller, tradingId, item.TradingId)
	}
	if item.Subject != subject {
		t.Errorf("%s Subject must be %s but %s", caller, subject, item.Subject)
	}
	if item.SortOrder != sortOrder {
		t.Errorf("%s SortOrder must be %d but %d", caller, sortOrder, item.SortOrder)
	}
	if item.UnitPrice != unitPrice {
		t.Errorf("%s UnitPrice must be %d but %d", caller, unitPrice, item.UnitPrice)
	}
	if item.Amount != amount {
		t.Errorf("%s Amount must be %f but %f", caller, amount, item.Amount)
	}
	if item.Degree != degree {
		t.Errorf("%s Degree must be %s but %s", caller, degree, item.Degree)
	}
	if item.TaxType != taxType {
		t.Errorf("%s TaxType must be %d but %d", caller, taxType, item.TaxType)
	}
	if item.Memo != memo {
		t.Errorf("%s Memo must be %s but %s", caller, memo, item.Memo)
	}
}
