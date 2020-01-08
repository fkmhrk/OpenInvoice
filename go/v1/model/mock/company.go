package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/company"
)

type CompanyDAO struct {
	GetListResult []*company.Company
	GetByIdResult *company.Company
	CreateResult  *company.Company
	UpdateResult  *company.Company
	DeleteResult  error
}

func (d *CompanyDAO) GetList() ([]*company.Company, error) {
	return d.GetListResult, nil
}

func (d *CompanyDAO) GetById(id string) (*company.Company, error) {
	return d.GetByIdResult, nil
}

func (d *CompanyDAO) Create(name, zip, address, phone, unit string) (*company.Company, error) {
	return d.CreateResult, nil
}

func (d *CompanyDAO) Update(id, name, zip, address, phone, unit string) (*company.Company, error) {
	return d.UpdateResult, nil
}

func (d *CompanyDAO) Delete(id string) error {
	return d.DeleteResult
}
