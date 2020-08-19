package company

import (
	"net/http"

	e "github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model/response"
	"github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"
	"github.com/mokelab-go/server/entity"
)

type companyService struct {
	sessionDAO model.Session
	companyDAO model.Company
}

// New creates instance
func New(models *model.Models) service.Company {
	return &companyService{
		sessionDAO: models.Session,
		companyDAO: models.Company,
	}
}

func (s *companyService) GetList() entity.Response {
	// get item
	companies, err := s.companyDAO.GetList()
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, c := range companies {
		list = append(list, s.toJson(c))
	}
	body := map[string]interface{}{
		"companies": list,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *companyService) Create(name, zip, address, phone, unit string) entity.Response {
	// create
	company, err := s.companyDAO.Create(name, zip, address, phone, unit)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": company.Id,
	}
	return entity.Response{
		Status: http.StatusCreated,
		Body:   body,
	}
}

func (s *companyService) Update(id, name, zip, address, phone, unit string) entity.Response {
	// create
	company, err := s.companyDAO.Update(id, name, zip, address, phone, unit)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": company.Id,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *companyService) Delete(id string) entity.Response {
	// delete
	err := o.companyDAO.Delete(id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	return entity.Response{
		Status: http.StatusNoContent,
	}
}

func (s *companyService) toJson(c *e.Company) map[string]interface{} {
	return map[string]interface{}{
		"id":      c.Id,
		"name":    c.Name,
		"zip":     c.Zip,
		"address": c.Address,
		"phone":   c.Phone,
		"unit":    c.Unit,
	}
}
