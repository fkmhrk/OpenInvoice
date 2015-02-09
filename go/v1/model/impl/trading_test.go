package impl

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func createTradingDAO(db *sql.DB) *tradingDAO {
	c := NewConnection(db)
	return NewTradingDAO(c)
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
		"work_from,work_to,quotation_date,bill_date," +
		"tax_rate,assignee,product,deleted)" +
		"VALUES(?,'company1',0,?,0,0,100,200,8.0,?,?,0)")
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
	date := "20150203"
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)
	deleteTradingId(db, date)

	dao := createTradingDAO(db)
	item, err := dao.Create(date, "company1111", "subject2222", 1,
		1234, 5678, quotationDate, billDate, taxRate, userId, "product3333")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	assertTrading(t, item, "20150203001", "company1111", "subject2222", 1,
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
	date := "20150203"
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)
	deleteTradingId(db, date)

	dao := createTradingDAO(db)
	item, err := dao.Create(date, "company1111", "subject2222", 1,
		1234, 5678, quotationDate, billDate, taxRate, userId, "product3333")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	if item.Id != "20150203001" {
		t.Errorf("Wrong ID : %s", item.Id)
	}

	// again
	item, err = dao.Create(date, "company4444", "subject5555", 1,
		1234, 5678, quotationDate, billDate, taxRate, userId, "product6666")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	if item.Id != "20150203002" {
		t.Errorf("Wrong ID : %s", item.Id)
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
	date := "20150203"
	userId := "user1122"
	quotationDate := int64(1111)
	billDate := int64(2222)
	taxRate := float32(8)
	deleteTradingByUser(db, userId)
	deleteTradingId(db, date)

	dao := createTradingDAO(db)
	item, err := dao.Create(date, "company1111", "subject2222", 1,
		1234, 5678, quotationDate, billDate, taxRate, userId, "product3333")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}

	// update
	item2, err := dao.Update(item.Id, "company2222", "subject3333",
		0, 2345, 6789, 2222, 3333, 10, userId, "product4444")
	if err != nil {
		t.Errorf("Failed to update trading : %s", err)
		return
	}
	if item2.Id != "20150203001" {
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
	if item3.Id != "20150203001" {
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
