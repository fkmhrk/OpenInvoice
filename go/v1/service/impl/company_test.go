package impl

import (
	"../../model"
	m "../../model/mock"
	"fmt"
	"testing"
)

func TestCompany0000_GetList(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	var list []*model.Company
	for i := 0; i < 3; i++ {
		list = append(list, &model.Company{
			Id:      fmt.Sprintf("company%d", i),
			Name:    fmt.Sprintf("Name%d", i),
			Zip:     fmt.Sprintf("Zip%d", i),
			Address: fmt.Sprintf("Address%d", i),
			Phone:   fmt.Sprintf("Phone%d", i),
			Unit:    fmt.Sprintf("Unit%d", i),
		})
	}
	companyDAO := &m.CompanyDAO{
		GetListResult: list,
	}

	s := NewCompanySerivce(sessionDAO, companyDAO)

	token := "token1122"
	r := s.GetList(token)
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
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	companyDAO := &m.CompanyDAO{
		CreateResult: &model.Company{
			Id:      "company",
			Name:    "Name",
			Zip:     "Zip",
			Address: "Address",
			Phone:   "Phone",
			Unit:    "Unit",
		},
	}

	s := NewCompanySerivce(sessionDAO, companyDAO)

	token := "token1122"
	r := s.Create(token, "name", "zip", "address", "phone", "unit")
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
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	companyDAO := &m.CompanyDAO{
		UpdateResult: &model.Company{
			Id:      "company",
			Name:    "Name",
			Zip:     "Zip",
			Address: "Address",
			Phone:   "Phone",
			Unit:    "Unit",
		},
	}

	s := NewCompanySerivce(sessionDAO, companyDAO)

	token := "token1122"
	r := s.Update(token, "id", "name", "zip", "address", "phone", "unit")
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
