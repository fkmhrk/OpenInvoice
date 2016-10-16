package impl

import (
	m "../../model"
	"../../model/mock"
	"fmt"
	"testing"
)

func TestCompany0000_GetList(t *testing.T) {
	models := mock.NewMock()
	var list []*m.Company
	for i := 0; i < 3; i++ {
		list = append(list, &m.Company{
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

	s := NewCompanySerivce(models)

	r := s.GetList()
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	companies, _ := json.Array("companies")
	if len(companies) != 3 {
		t.Errorf("Wrong length : %d", len(companies))
		return
	}
	company, _ := companies.Object(0)
	assertCompany(t, company, "company0", "Name0", "Zip0", "Address0", "Phone0", "Unit0")
	company, _ = companies.Object(1)
	assertCompany(t, company, "company1", "Name1", "Zip1", "Address1", "Phone1", "Unit1")
	company, _ = companies.Object(2)
	assertCompany(t, company, "company2", "Name2", "Zip2", "Address2", "Phone2", "Unit2")
}

func TestCompany0100_Create(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.CreateResult = &m.Company{
		Id:      "company",
		Name:    "Name",
		Zip:     "Zip",
		Address: "Address",
		Phone:   "Phone",
		Unit:    "Unit",
	}

	s := NewCompanySerivce(models)

	r := s.Create("name", "zip", "address", "phone", "unit")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 201 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "company" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestCompany0200_Update(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.UpdateResult = &m.Company{
		Id:      "company",
		Name:    "Name",
		Zip:     "Zip",
		Address: "Address",
		Phone:   "Phone",
		Unit:    "Unit",
	}

	s := NewCompanySerivce(models)

	r := s.Update("id", "name", "zip", "address", "phone", "unit")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "company" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestCompany0300_Delete(t *testing.T) {
	models := mock.NewMock()
	companyDAO, _ := models.Company.(*mock.CompanyDAO)
	companyDAO.DeleteResult = nil

	s := NewCompanySerivce(models)

	result := s.Delete("company1")
	if result == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if result.Status() != 204 {
		t.Errorf("Wrong status : %d", result.Status())
		return
	}
	if result.Body() != "" {
		t.Errorf("Body must be empty but %s", result.Body())
		return
	}
}
