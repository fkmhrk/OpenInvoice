package impl

import (
	s "../"
	m "../../model"
)

type companyService struct {
	sessionDAO m.SessionDAO
	companyDAO m.CompanyDAO
}

func NewCompanySerivce(s m.SessionDAO, c m.CompanyDAO) *companyService {
	return &companyService{
		sessionDAO: s,
		companyDAO: c,
	}
}

func (s *companyService) GetList(token string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get item
	companies, err := s.companyDAO.GetList()
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, c := range companies {
		list = append(list, s.toJson(c))
	}
	body := map[string]interface{}{
		"companies": list,
	}
	return jsonResult(200, body)
}

func (s *companyService) Create(token, name, zip, address, phone, unit string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// create
	company, err := s.companyDAO.Create(name, zip, address, phone, unit)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": company.Id,
	}
	return jsonResult(201, body)
}

func (s *companyService) Update(token, id, name, zip, address, phone, unit string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// create
	company, err := s.companyDAO.Update(id, name, zip, address, phone, unit)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": company.Id,
	}
	return jsonResult(200, body)
}

func (s *companyService) toJson(c *m.Company) map[string]interface{} {
	return map[string]interface{}{
		"id":      c.Id,
		"name":    c.Name,
		"zip":     c.Zip,
		"address": c.Address,
		"phone":   c.Phone,
		"unit":    c.Unit,
	}
}
