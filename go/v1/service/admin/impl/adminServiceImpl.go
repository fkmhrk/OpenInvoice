package impl

import (
	"net/http"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
	"github.com/fkmhrk/OpenInvoice/v1/model/response"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/mokelab-go/server/entity"
)

type adminService struct {
	sessionDAO session.SessionDAO
	envDAO     env.DAO
}

// New creates instance
func New(models *m.Models) *adminService {
	return &adminService{
		sessionDAO: models.Session,
		envDAO:     models.Env,
	}
}

func (o *adminService) GetEnvironment() entity.Response {
	// get list
	envList, err := o.envDAO.GetList()
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := make(map[string]interface{}, 0)
	for _, env := range envList {
		body[env.Key] = env.Value
	}

	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *adminService) SaveEnvironment(list []*env.Env) entity.Response {
	// saves
	err := o.envDAO.Save(list)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"msg": "ok",
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *adminService) GetMyCompanyname() entity.Response {
	// get list
	env, err := o.envDAO.Get("company_name")
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	var name string
	if env.IsEmpty() {
		name = ""
	} else {
		name = env.Value
	}

	body := map[string]interface{}{
		"name": name,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}
