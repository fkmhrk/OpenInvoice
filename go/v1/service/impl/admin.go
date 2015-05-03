package impl

import (
	s "../"
	m "../../model"
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

func (o *adminService) GetEnvironment(token string) s.Result {
	if isEmpty(token) {
		return errorResult(400, s.ERR_AUTHORIZATION_EMPTY)
	}
	// get Session
	session, err := o.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, s.ERR_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, s.ERR_NOT_AUTHORIZED)
	}
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
