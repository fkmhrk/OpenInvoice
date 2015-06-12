package impl

import (
	s "../"
	m "../../model"
	"../../model/mock"
	"testing"
)

func TestAdmin_0000_GetEnvironment(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token:  "token1122",
		UserId: "user1122",
	}
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.GetListResult = []*m.Env{
		&m.Env{
			Key:   "company_name",
			Value: "MyCompany1",
		},
		&m.Env{
			Key:   "company_tel",
			Value: "080-1111-2222",
		},
	}

	var service s.AdminService = NewAdminService(models)
	r := service.GetEnvironment("token")

	// Assertion
	if r.Status() != 200 {
		t.Errorf("Status must be 200 but %d", r.Status())
		return
	}
	json := json(r)
	if len(json) != 2 {
		t.Errorf("JSON length must be 2 but %d", len(json))
		return
	}
	assertString(t, json, "company_name", "MyCompany1")
	assertString(t, json, "company_tel", "080-1111-2222")
}

func TestAdmin_0200_SavetEnvironment(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token:  "token1122",
		UserId: "user1122",
	}
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.SaveResult = nil

	list := []*m.Env{
		&m.Env{
			Key:   "company_name",
			Value: "mokelab inc",
		},
	}

	var service s.AdminService = NewAdminService(models)
	r := service.SaveEnvironment("token", list)

	// Assertion
	if r.Status() != 200 {
		t.Errorf("Status must be 200 but %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "msg", "ok")
}

func TestAdmin_0400_GetMyCompanyname(t *testing.T) {
	models := mock.NewMock()
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.GetResult = m.Env{
		Key:   "company_name",
		Value: "mokelab inc",
	}

	var service s.AdminService = NewAdminService(models)
	r := service.GetMyCompanyname()

	// Assertion
	if r.Status() != 200 {
		t.Errorf("Status must be 200 but %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "name", "mokelab inc")
}
