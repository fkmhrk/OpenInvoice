package company

import (
	"fmt"
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model/mock"
	"github.com/fkmhrk/OpenInvoice/v1/model/test"
)

func TestCompany0000_GetList(t *testing.T) {
	models := mock.NewMock()
	var list []*entity.Company
	for i := 0; i < 3; i++ {
		list = append(list, &entity.Company{
			Id:      fmt.Sprintf("company%d", i),
			Name:    fmt.Sprintf("Name%d", i),
			Zip:     fmt.Sprintf("Zip%d", i),
			Address: fmt.Sprintf("Address%d", i),
			Phone:   fmt.Sprintf("Phone%d", i),
			Unit:    fmt.Sprintf("Unit%d", i),
		})
	}
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.GetListResult = list

	s := New(models)

	r := s.GetList()
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	companies, _ := r.Body["companies"].([]interface{})
	if len(companies) != 3 {
		t.Errorf("Wrong length : %d", len(companies))
		return
	}
	company, _ := companies[0].(map[string]interface{})
	assertCompany(t, company, "company0", "Name0", "Zip0", "Address0", "Phone0", "Unit0")
	company, _ = companies[1].(map[string]interface{})
	assertCompany(t, company, "company1", "Name1", "Zip1", "Address1", "Phone1", "Unit1")
	company, _ = companies[2].(map[string]interface{})
	assertCompany(t, company, "company2", "Name2", "Zip2", "Address2", "Phone2", "Unit2")
}

func TestCompany0100_Create(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.CreateResult = &entity.Company{
		Id:      "company",
		Name:    "Name",
		Zip:     "Zip",
		Address: "Address",
		Phone:   "Phone",
		Unit:    "Unit",
	}

	s := New(models)

	r := s.Create("name", "zip", "address", "phone", "unit")
	if r.Status != 201 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "company" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestCompany0200_Update(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.UpdateResult = &entity.Company{
		Id:      "company",
		Name:    "Name",
		Zip:     "Zip",
		Address: "Address",
		Phone:   "Phone",
		Unit:    "Unit",
	}

	s := New(models)

	r := s.Update("id", "name", "zip", "address", "phone", "unit")
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "company" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestCompany0300_Delete(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.DeleteResult = nil

	s := New(models)

	result := s.Delete("company1")
	if result.Status != 204 {
		t.Errorf("Wrong status : %d", result.Status)
		return
	}
	if result.Body != nil {
		t.Errorf("Body must be empty but %s", result.Body)
		return
	}
}

func assertCompany(t *testing.T, item map[string]interface{}, id, name, zip, address, phone, unit string) bool {
	if test.AssertString(t, item, "id", id) ||
		test.AssertString(t, item, "name", name) ||
		test.AssertString(t, item, "zip", zip) ||
		test.AssertString(t, item, "address", address) ||
		test.AssertString(t, item, "phone", phone) ||
		test.AssertString(t, item, "unit", unit) {
		return true
	}
	return false
}
