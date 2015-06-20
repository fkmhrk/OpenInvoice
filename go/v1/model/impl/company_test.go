package impl

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func createCompanyDAO(db *sql.DB) *companyDAO {
	c := NewConnection(db)
	return NewCompanyDAO(c)
}

func deleteCompanyByName(db *sql.DB, name string) {
	s, _ := db.Prepare("DELETE FROM company WHERE name=?")
	defer s.Close()
	s.Exec(name)
}

func insertCompany(db *sql.DB, id, name, zip, address, phone, unit string) {
	s, _ := db.Prepare("INSERT INTO company(" +
		"id,name,zip,address,phone,unit,deleted)" +
		"VALUES(?,?,?,?,?,?,0)")
	defer s.Close()
	s.Exec(id, name, zip, address, phone, unit)
}

func TestCompany0000_GetList(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	name := "name1"
	deleteCompanyByName(db, name)

	dao := createCompanyDAO(db)
	list, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to get company list : %s", err)
		return
	}
	// insert
	insertCompany(db, "id1", name, "zip1", "address1", "phone1", "unit1")
	list2, err := dao.GetList()

	if len(list2) != len(list)+1 {
		t.Errorf("Wrong list length : %d", len(list2))
		return
	}
}

func TestCompany0100_GetById(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	name := "name1"
	deleteCompanyByName(db, name)
	insertCompany(db, "id1", name, "zip1", "address1", "phone1", "unit1")
	dao := createCompanyDAO(db)
	item, err := dao.GetById("id1")
	if err != nil {
		t.Errorf("Failed to get company by id : %s", err)
		return
	}
	if item == nil {
		t.Errorf("Failed to get company by id : nil")
		return
	}
	assertCompany(t, item, "id1", name, "zip1", "address1", "phone1", "unit1")
}

func TestCompany0200_Create(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	name := "name1"
	zip := "zip2"
	address := "address2"
	phone := "phone2"
	unit := "unit2"
	deleteCompanyByName(db, name)

	dao := createCompanyDAO(db)
	item, err := dao.Create(name, zip, address, phone, unit)
	if err != nil {
		t.Errorf("Failed to create company : %s", err)
		return
	}
	assertCompany(t, item, item.Id, name, zip, address, phone, unit)
	// get by id
	item2, err := dao.GetById(item.Id)
	if err != nil {
		t.Errorf("Failed to created company : %s", err)
		return
	}
	assertCompany(t, item2, item.Id, name, zip, address, phone, unit)
}

func TestCompany0300_Update(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	name := "name1"
	zip := "zip2"
	address := "address2"
	phone := "phone2"
	unit := "unit2"
	deleteCompanyByName(db, name)

	dao := createCompanyDAO(db)
	item, err := dao.Create(name, zip, address, phone, unit)
	if err != nil {
		t.Errorf("Failed to create company : %s", err)
		return
	}
	assertCompany(t, item, item.Id, name, zip, address, phone, unit)
	// update
	item2, err := dao.Update(item.Id, "name3", "zip3", "address3", "phone3", "unit3")
	if err != nil {
		t.Errorf("Failed to update company : %s", err)
		return
	}
	assertCompany(t, item2, item.Id, "name3", "zip3", "address3", "phone3", "unit3")
	// get by id
	item3, err := dao.GetById(item.Id)
	if err != nil {
		t.Errorf("Failed to updated company : %s", err)
		return
	}
	assertCompany(t, item3, item.Id, "name3", "zip3", "address3", "phone3", "unit3")
}

func TestCompany0400_Delete(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	name := "name1"
	zip := "zip2"
	address := "address2"
	phone := "phone2"
	unit := "unit2"
	userId := "user1"
	deleteCompanyByName(db, name)
	deleteTradingByUser(db, userId)

	dao := createCompanyDAO(db)
	item, err := dao.Create(name, zip, address, phone, unit)
	if err != nil {
		t.Errorf("Failed to create company : %s", err)
		return
	}
	assertCompany(t, item, item.Id, name, zip, address, phone, unit)
	tradingDAO := createTradingDAO(db)
	trading, err := tradingDAO.Create(item.Id, "subject", 1, 100, 200, 1280,
		1111, 2222, 3333, 8, userId, "product", "memo")
	if err != nil {
		t.Errorf("Failed to create trading")
		return
	}

	// delete
	err = dao.Delete(item.Id)
	if err != nil {
		t.Errorf("Failed to delete company : %s", err)
		return
	}

	// check
	company, err := dao.GetById(item.Id)
	if err != nil {
		t.Errorf("Failed to get company : %s", err)
		return
	}
	if company != nil {
		t.Errorf("Unexpected get company result")
		return
	}

	// trading
	trading2, err := tradingDAO.GetById(trading.Id, userId)
	if err != nil {
		t.Errorf("Failed to get trading by Id : %s", err)
		return
	}
	if trading2.CompanyId != "" {
		t.Errorf("Company ID must be empty but %s", trading2.CompanyId)
		return
	}
}
