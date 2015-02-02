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
