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
		"id,company_id,subject," +
		"work_from,work_to,assignee,product,deleted)" +
		"VALUES(?,'company1',?,0,0,?,?,0)")
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
	item := list[0]
	if item.Id != "trade1" {
		t.Errorf("Wrong ID : %s", item.Id)
	}
	if item.Subject != "subject1" {
		t.Errorf("Wrong Subject : %s", item.Subject)
	}
	if item.Product != "product2233" {
		t.Errorf("Wrong Product : %s", item.Product)
	}
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
	deleteTradingByUser(db, userId)
	deleteTradingId(db, date)

	dao := createTradingDAO(db)
	item, err := dao.Create(date, "company1111", "subject2222",
		1234, 5678, userId, "product3333")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	if item.Id != "20150203001" {
		t.Errorf("Wrong ID : %s", item.Id)
	}
	if item.CompanyId != "company1111" {
		t.Errorf("Wrong Company ID : %s", item.CompanyId)
	}
	if item.Subject != "subject2222" {
		t.Errorf("Wrong Subject : %s", item.Subject)
	}
	if item.Product != "product3333" {
		t.Errorf("Wrong Product : %s", item.Product)
	}
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
	deleteTradingByUser(db, userId)
	deleteTradingId(db, date)

	dao := createTradingDAO(db)
	item, err := dao.Create(date, "company1111", "subject2222",
		1234, 5678, userId, "product3333")
	if err != nil {
		t.Errorf("Failed to create tradings : %s", err)
		return
	}
	if item.Id != "20150203001" {
		t.Errorf("Wrong ID : %s", item.Id)
	}

	// again
	item, err = dao.Create(date, "company4444", "subject5555",
		1234, 5678, userId, "product6666")
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
