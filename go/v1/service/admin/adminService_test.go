package admin

import (
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model/mock"
	"github.com/fkmhrk/OpenInvoice/v1/model/test"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
)

func TestAdmin_0000_GetEnvironment(t *testing.T) {
	models := mock.NewMock()
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.GetListResult = []*entity.Env{
		&entity.Env{
			Key:   "company_name",
			Value: "MyCompany1",
		},
		&entity.Env{
			Key:   "company_tel",
			Value: "080-1111-2222",
		},
	}

	var service service.Admin = New(models)
	r := service.GetEnvironment()

	// Assertion
	if r.Status != 200 {
		t.Errorf("Status must be 200 but %d", r.Status)
		return
	}
	if len(r.Body) != 2 {
		t.Errorf("JSON length must be 2 but %d", len(r.Body))
		return
	}
	test.AssertString(t, r.Body, "company_name", "MyCompany1")
	test.AssertString(t, r.Body, "company_tel", "080-1111-2222")
}

func TestAdmin_0200_SavetEnvironment(t *testing.T) {
	models := mock.NewMock()
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.SaveResult = nil

	list := []*entity.Env{
		&entity.Env{
			Key:   "company_name",
			Value: "mokelab inc",
		},
	}

	var service service.Admin = New(models)
	r := service.SaveEnvironment(list)

	// Assertion
	if r.Status != 200 {
		t.Errorf("Status must be 200 but %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "msg", "ok")
}

func TestAdmin_0400_GetMyCompanyname(t *testing.T) {
	models := mock.NewMock()
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.GetResult = entity.Env{
		Key:   "company_name",
		Value: "mokelab inc",
	}

	var service service.Admin = New(models)
	r := service.GetMyCompanyname()

	// Assertion
	if r.Status != 200 {
		t.Errorf("Status must be 200 but %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "name", "mokelab inc")
}
