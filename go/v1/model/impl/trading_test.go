package impl

import (
	m "../"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func createTradingDAO(db *sql.DB) *tradingDAO {
	c := NewConnection(db)
	return NewTradingDAO(c, NewLogger())
}

func deleteTradingByUser(db *sql.DB, userId string) {
	s, _ := db.Prepare("DELETE FROM trading WHERE assignee=?")
	defer s.Close()
	s.Exec(userId)
}

func deleteTradingId(db *sql.DB, date string) {
	s, _ := db.Prepare("DELETE FROM trading_id WHERE date=?")
	defer s.Close()
	s.Exec(date)
}

func insertTrading(db *sql.DB, id, user, subject, product string) {
	s, _ := db.Prepare("INSERT INTO trading(" +
		"id,company_id,title_type,subject," +
		"work_from,work_to,total," +
		"quotation_date,quotation_number," +
		"bill_date,bill_number," +
		"delivery_date,delivery_number," +
		"tax_rate,assignee,product,memo,created_time,modified_time,deleted)" +
		"VALUES(?,'company1',0,?," +
		"0,0,1280," +
		"100,''," +
		"200,''," +
		"300,''," +
		"8.0,?,?,'memo',unix_timestamp(now()),unix_timestamp(now()),0)")
	defer s.Close()
	s.Exec(id, subject, user, product)
}

func TestTrading0000_GetListByUser(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	product := "product2233"
	deleteTradingByUser(db, userId)
	insertTrading(db, "trade1", userId, "subject1", product)
	insertTrading(db, "trade2", userId, "subject2", product)

	dao := createTradingDAO(db)
	list, err := dao.GetListByUser(userId)
	if err != nil {
		t.Errorf("Failed to get tradings by name : %s", err)
		return
	}
	if len(list) != 2 {
		t.Errorf("Wrong list length : %d", len(list))
		return
	}
	assertTrading(t, list[0], "trade1", "company1", "subject1", 0,
		0, 0, 100, 200, 8, "user1122", "product2233")
	assertTrading(t, list[1], "trade2", "company1", "subject2", 0,
		0, 0, 100, 200, 8, "user1122", "product2233")
}

func TestTrading0001_GetListByUser_0(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	deleteTradingByUser(db, userId)

	dao := createTradingDAO(db)
	list, err := dao.GetListByUser(userId)
	if err != nil {
		t.Errorf("Failed to get tradings by name : %s", err)
		return
	}
	if len(list) != 0 {
		t.Errorf("Wrong list length : %d", len(list))
		return
	}
}

func TestTrading0100_GetById(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	product := "product2233"
	deleteTradingByUser(db, userId)
	insertTrading(db, "trade1", userId, "subject1", product)
	insertTrading(db, "trade2", userId, "subject2", product)

	dao := createTradingDAO(db)
	item, err := dao.GetById("trade1", userId)
	if err != nil {
		t.Errorf("Failed to get tradings by Id : %s", err)
		return
	}
	assertTrading(t, item, "trade1", "company1", "subject1", 0,
		0, 0, 100, 200, 8, "user1122", "product2233")
}

func TestTrading0101_GetById_noId(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	product := "product2233"
	deleteTradingByUser(db, userId)
	insertTrading(db, "trade1", userId, "subject1", product)
	insertTrading(db, "trade2", userId, "subject2", product)

	dao := createTradingDAO(db)
	item, err := dao.GetById("trade3", userId)
	if err != nil {
		t.Errorf("Failed to get tradings by Id : %s", err)
		return
	}
	if item != nil {
		t.Errorf("item must be nil ID=%s", item.Id)
	}
}

func TestTrading0100_Create(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	deliveryDate := int64(3333)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)

	dao := createTradingDAO(db)
	item, err := dao.Create("company1111", "subject2222", 1,
		1234, 5678, 1280, quotationDate, billDate, deliveryDate, taxRate, userId, "product3333", "memo")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	assertTrading(t, item, item.Id, "company1111", "subject2222", 1,
		1234, 5678, 1111, 2222, 8, "user1122", "product3333")
}

func TestTrading0101_Create_2(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	deliveryDate := int64(3333)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)

	dao := createTradingDAO(db)
	item, err := dao.Create("company1111", "subject2222", 1,
		1234, 5678, 2980, quotationDate, billDate, deliveryDate, taxRate, userId, "product3333", "memo")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}

	// again
	item, err = dao.Create("company4444", "subject5555", 1,
		1234, 5678, 3980, quotationDate, billDate, deliveryDate, taxRate, userId, "product6666", "memo")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	if item.CompanyId != "company4444" {
		t.Errorf("Wrong Company ID : %s", item.CompanyId)
	}
	if item.Subject != "subject5555" {
		t.Errorf("Wrong Subject : %s", item.Subject)
	}
	if item.Product != "product6666" {
		t.Errorf("Wrong Product : %s", item.Product)
	}
}

func TestTrading0300_Update(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	deliveryDate := int64(3333)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)

	dao := createTradingDAO(db)
	item, err := dao.Create("company1111", "subject2222", 1,
		1234, 5678, 4980, quotationDate, billDate, deliveryDate, taxRate, userId, "product3333", "memo")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}

	// update
	var total int64 = 3333
	item2, err := dao.Update(m.Trading{
		Id:            item.Id,
		CompanyId:     "company2222",
		Subject:       "subject3333",
		TitleType:     0,
		WorkFrom:      2345,
		WorkTo:        6789,
		Total:         total,
		QuotationDate: 2222,
		BillDate:      3333,
		TaxRate:       10,
		AssigneeId:    userId,
		Product:       "product4444",
	})
	if err != nil {
		t.Errorf("Failed to update trading : %s", err)
		return
	}
	if item2.Id != item.Id {
		t.Errorf("Wrong ID : %s", item2.Id)
	}
	if item2.CompanyId != "company2222" {
		t.Errorf("Wrong Company ID : %s", item2.CompanyId)
	}
	if item2.Subject != "subject3333" {
		t.Errorf("Wrong Subject : %s", item2.Subject)
	}
	if item2.Product != "product4444" {
		t.Errorf("Wrong Product : %s", item2.Product)
	}

	// get by id
	item3, err := dao.GetById(item.Id, userId)
	if err != nil {
		t.Errorf("Failed to get trading : %s", err)
		return
	}
	if item3.Id != item.Id {
		t.Errorf("Wrong ID : %s", item3.Id)
	}
	if item3.CompanyId != "company2222" {
		t.Errorf("Wrong Company ID : %s", item3.CompanyId)
	}
	if item3.Subject != "subject3333" {
		t.Errorf("Wrong Subject : %s", item3.Subject)
	}
	if item3.Product != "product4444" {
		t.Errorf("Wrong Product : %s", item3.Product)
	}
}

func TestTrading0400_Delete(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	deliveryDate := int64(3333)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)
	// create trading
	dao := createTradingDAO(db)
	trading, err := dao.Create("company1111", "subject2222", 1,
		1234, 5678, 4980, quotationDate, billDate, deliveryDate, taxRate, userId, "product3333", "memo")
	if err != nil {
		t.Errorf("Failed to create trading : %s", err)
		return
	}
	// add item
	for i := 0; i < 2; i++ {
		_, err = dao.CreateItem(trading.Id, "item", "m/m", "memo", i, 100, i, 1)
		if err != nil {
			t.Errorf("Failed to create trading item : %s", err)
			return
		}
	}
	// check
	itemList, err := dao.GetItemsById(trading.Id)
	if err != nil {
		t.Errorf("Failed to get items : %s", err)
		return
	}
	if len(itemList) != 2 {
		t.Errorf("ItemList must be 2 but %d", len(itemList))
		return
	}
	// delete
	err = dao.Delete(trading.Id)
	if err != nil {
		t.Errorf("Failed to delete trading : %s", err)
		return
	}
	// check
	itemList2, err := dao.GetItemsById(trading.Id)
	if err != nil {
		t.Errorf("Failed to get items : %s", err)
		return
	}
	if len(itemList2) != 0 {
		t.Errorf("ItemList must be 0 but %d", len(itemList2))
		return
	}
	trading2, err := dao.GetById(trading.Id, userId)
	if err != nil {
		t.Errorf("Failed to get trading : %s", err)
		return
	}
	if trading2 != nil {
		t.Errorf("Trading is not deleted...")
		return
	}
}
