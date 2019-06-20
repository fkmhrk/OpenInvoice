package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
)

type adminService struct {
	sessionDAO m.SessionDAO
	envDAO     m.EnvDAO
}

func NewAdminService(models *m.Models) *adminService {
	return &adminService{
		sessionDAO: models.Session,
		envDAO:     models.Env,
	}
}

func (o *adminService) GetEnvironment() s.Result {
	// get list
	envList, err := o.envDAO.GetList()
	if err != nil {
		return errorResult(500, s.ERR_SERVER_ERROR)
	}

	body := make(map[string]interface{}, 0)
	for _, env := range envList {
		body[env.Key] = env.Value
	}

	return jsonResult(200, body)
}

func (o *adminService) SaveEnvironment(list []*m.Env) s.Result {
	// saves
	err := o.envDAO.Save(list)
	if err != nil {
		return errorResult(500, s.ERR_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"msg": "ok",
	}
	return jsonResult(200, body)
}

func (o *adminService) GetMyCompanyname() s.Result {
	// get list
	env, err := o.envDAO.Get("company_name")
	if err != nil {
		return errorResult(500, s.ERR_SERVER_ERROR)
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
	return jsonResult(200, body)
}
