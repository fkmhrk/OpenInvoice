package mock

import (
	m "../"
)

type CompanyDAO struct {
	GetListResult []*m.Company
	GetByIdResult *m.Company
	CreateResult  *m.Company
	UpdateResult  *m.Company
	DeleteResult  error
}

func (d *CompanyDAO) GetList() ([]*m.Company, error) {
	return d.GetListResult, nil
}

func (d *CompanyDAO) GetById(id string) (*m.Company, error) {
	return d.GetByIdResult, nil
}

func (d *CompanyDAO) Create(name, zip, address, phone, unit string) (*m.Company, error) {
	return d.CreateResult, nil
}

func (d *CompanyDAO) Update(id, name, zip, address, phone, unit string) (*m.Company, error) {
	return d.UpdateResult, nil
}

func (d *CompanyDAO) Delete(id string) error {
	return d.DeleteResult
}
